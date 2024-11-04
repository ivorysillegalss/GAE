package crawl

import (
	"context"
	"encoding/json"
	"gae-backend-crawl/bootstrap"
	"gae-backend-crawl/constant/mq"
	"gae-backend-crawl/constant/sys"
	"gae-backend-crawl/domain"
	"gae-backend-crawl/infrastructure/bloom"
	kq "gae-backend-crawl/infrastructure/kafka"
	"gae-backend-crawl/infrastructure/log"
	"github.com/google/go-github/github"
	_ "github.com/json-iterator/go"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type repoCrawl struct {
	bloom     bloom.Client
	env       *bootstrap.Env
	kafkaConf *bootstrap.KafkaConf
}

func getRepoId() int64 {
	return atomic.AddInt64(&repoId, 1)
}

var (
	repoId            int64
	userMessagePusher *kq.Pusher
	repoMessagePusher *kq.Pusher
	ctx               context.Context
	client            *github.Client
	tokenManager      *TokenManager
)

func init() {
	ctx = context.Background()
	repoId = 675
}

// 注册MQ相关队列
func registerMessageQueue(r *repoCrawl) {
	kqUserConf := r.kafkaConf.Conf[mq.UnCleansingUserId]
	userMessagePusher = kq.NewPusher(kqUserConf.Brokers, kqUserConf.Topic)

	kqRepoConf := r.kafkaConf.Conf[mq.UnCleansingRepoId]
	repoMessagePusher = kq.NewPusher(kqRepoConf.Brokers, kqRepoConf.Topic)
}

func crawlPushContributorsInfo(contributors *[]*github.Contributor, b bloom.Client) *[]string {
	var contributorsValue []*domain.Contributor
	var contributorsId []string
	for _, contributor := range *contributors {

		checkSwitchToken()

		userV, _, err := client.Users.GetByID(ctx, *contributor.ID)

		if userV == nil {
			continue
		}

		formatId := strconv.FormatInt(*userV.ID, 10)
		if err != nil {
			log.GetTextLogger().Warn("error getting userInfo for user: ")
			continue
		}

		contributorsId = append(contributorsId, formatId)

		v := domain.NewContributorValue(contributor, userV)

		follower, i := crawlContributorFollower(userV.GetLogin())
		if i == -1 {
			return nil
		} else {
			v.Followers = *follower
		}
		following, j := crawlContributorFollowing(userV.GetLogin())
		if j == -1 {
			return nil
		} else {
			v.Followings = *following
		}

		contributorsValue = append(contributorsValue, v)

		bloomCheckBeforePush(v, b)
	}
	return &contributorsId
}

func crawlContributorFollower(username string) (*[]*github.User, int) {
	checkSwitchToken()

	var allFollowers []*github.User
	opts := &github.ListOptions{PerPage: 100}
	followers, _, err := client.Users.ListFollowers(ctx, username, opts)
	if followers == nil {
		return nil, -1
	}
	if err != nil {
		log.GetTextLogger().Warn("list followers error for user: " + username)
	}
	allFollowers = append(allFollowers, followers...)
	return &allFollowers, 0
}

func crawlContributorFollowing(username string) (*[]*github.User, int) {
	checkSwitchToken()

	var allFollowings []*github.User
	opts := &github.ListOptions{PerPage: 100}
	following, _, err := client.Users.ListFollowing(ctx, username, opts)
	if following == nil {
		return nil, -1
	}
	if err != nil {
		log.GetTextLogger().Warn("list followers error for user: " + username)
	}
	allFollowings = append(allFollowings, following...)
	return &allFollowings, 0
}

func bloomCheckBeforePush(v *domain.Contributor, b bloom.Client) {
	formatId := strconv.FormatInt(v.Id, 10)
	isExist := b.Check(formatId)
	if !isExist {
		//log.GetTextLogger().Info("success user: ", v)

		marshal, _ := json.Marshal(v)
		err := userMessagePusher.KPush(ctx, strconv.Itoa(int(time.Now().Unix())), string(marshal))
		if err != nil {
			log.GetTextLogger().Error("error pushing msg,err: ", err.Error())
		}

		b.Add(formatId)
	}
}

// DoCrawl 爬取仓库的主方法
func (r *repoCrawl) DoCrawl() {
	err := r.bloom.LoadFromFile(sys.BloomFilterFileName)
	if err != nil {
		log.GetTextLogger().Warn("Error loading bloom filter:", err)
	} else {
		log.GetTextLogger().Info("Bloom filter data loaded or new file created.")
	}

	r.bloom.StartAutoSave(sys.BloomFilterFileName, sys.BloomFlushDuration)

	registerMessageQueue(r)
	tokenValue := strings.Split(r.env.GithubTokens, ",")
	tokenManager = NewTokenManager(tokenValue)

	var wg sync.WaitGroup
	for i := 0; i < mq.MaxGoroutine; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {

				//如果这里没有可用的 会阻塞住
				client = tokenManager.GetClient(ctx)

				// 检查限额
				checkSwitchToken()

				repoId := getRepoId()
				repo, _, err := client.Repositories.GetByID(ctx, repoId)
				if err != nil || repo == nil {
					log.GetTextLogger().Info("Fetching nil repository by Id: %v", err)
					continue
				}

				// 仓库信息爬取完成，进行初步清洗
				value := domain.NewRepositoryValue(repo)

				// 获取该项目的所有贡献者
				var allContributors []*github.Contributor

				for {
					checkSwitchToken()

					opts := &github.ListContributorsOptions{ListOptions: github.ListOptions{PerPage: 100}}
					contributors, resp, err := client.Repositories.ListContributors(ctx, *repo.Owner.Login, *repo.Name, opts)
					if contributors == nil {
						continue
					}

					if err != nil {
						log.GetTextLogger().Error("error fetching contributors: %v", err)
					}

					checkSwitchToken()

					allContributors = append(allContributors, contributors...)

					if resp == nil {
						continue
					}

					if resp.NextPage == 0 {
						break
					}
					opts.Page = resp.NextPage
				}

				contributorsInfo := crawlPushContributorsInfo(&allContributors, r.bloom)
				if contributorsInfo == nil {
					continue
				}
				value.ContributorsId = contributorsInfo

				log.GetTextLogger().Info("success repo: ", strconv.FormatInt(repoId, 10))
				marshal, _ := jsoniter.Marshal(value)
				err = repoMessagePusher.KPush(ctx, strconv.Itoa(int(time.Now().Unix())), string(marshal))

				if err != nil {
					log.GetTextLogger().Error("error pushing msg,err: ", err.Error())
				}

				time.Sleep(2 * time.Second)
			}
		}()
	}
	wg.Wait()
}

func checkSwitchToken() {
	if tokenManager.CheckRateLimit(client) {
		log.GetTextLogger().Warn("Token rate limit reached, switching...")
		client = tokenManager.GetClient(ctx)
	}
}

func NewRepoCrawl(b bloom.Client, c *bootstrap.KafkaConf, env *bootstrap.Env) domain.RepoCrawl {
	return &repoCrawl{bloom: b, kafkaConf: c, env: env}
}
