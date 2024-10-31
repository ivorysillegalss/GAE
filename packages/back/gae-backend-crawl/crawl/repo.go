package crawl

import (
	"context"
	"encoding/json"
	"gae-backend-crawl/bootstrap"
	"gae-backend-crawl/constant/mq"
	"gae-backend-crawl/domain"
	"gae-backend-crawl/infrastructure/bloom"
	kq "gae-backend-crawl/infrastructure/kafka"
	"gae-backend-crawl/infrastructure/log"
	"github.com/google/go-github/github"
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
)

func init() {
	ctx = context.Background()
}

// 注册MQ相关队列
func registerMessageQueue(r *repoCrawl) {
	kqUserConf := r.kafkaConf.Conf[mq.UnCleansingUserId]
	userMessagePusher = kq.NewPusher(kqUserConf.Brokers, kqUserConf.Topic)

	kqRepoConf := r.kafkaConf.Conf[mq.UnCleansingRepoId]
	repoMessagePusher = kq.NewPusher(kqRepoConf.Brokers, kqRepoConf.Topic)
}

// DoCrawl 爬取仓库的主方法
func (r *repoCrawl) DoCrawl() {
	registerMessageQueue(r)
	tokenValue := strings.Split(r.env.GithubTokens, ",")
	tokenManager := NewTokenManager(tokenValue)

	var wg sync.WaitGroup
	for i := 0; i < mq.MaxGoroutine; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				client := tokenManager.GetClient(ctx)
				// 检查限额
				if tokenManager.CheckRateLimit(client) {
					log.GetTextLogger().Warn("Token rate limit reached, switching...")
					continue
				}

				repoId := getRepoId()
				repo, _, err := client.Repositories.GetByID(ctx, repoId)
				if err != nil {
					log.GetTextLogger().Error("Error fetching repository by Id: %v", err)
					continue
				}

				// 获取该项目的所有贡献者
				var allContributors []*github.Contributor
				for {
					if tokenManager.CheckRateLimit(client) {
						log.GetTextLogger().Warn("Token rate limit reached, switching...")
						continue
					}

					opts := &github.ListContributorsOptions{ListOptions: github.ListOptions{PerPage: 100}}
					contributors, resp, err := client.Repositories.ListContributors(ctx, *repo.Owner.Login, *repo.Name, opts)

					if err != nil {
						log.GetTextLogger().Error("error fetching contributors: %v", err)
					}
					allContributors = append(allContributors, contributors...)

					if resp.NextPage == 0 {
						break
					}
					opts.Page = resp.NextPage
				}

				// 仓库信息与贡献者信息爬取完成，进行初步清洗
				value := domain.NewRepositoryValue(repo)
				for _, contributor := range allContributors {
					formatId := strconv.FormatInt(*contributor.ID, 10)
					id := *value.ContributorsId
					id = append(id, formatId)
					value.ContributorsId = &id

					isExist := r.bloom.Check(formatId)
					if !isExist {
						contributorValue := domain.NewContributorValue(contributor)
						marshal, _ := json.Marshal(contributorValue)

						err := userMessagePusher.KPush(ctx, strconv.Itoa(int(time.Now().Unix())), string(marshal))
						if err != nil {
							log.GetTextLogger().Error("error pushing msg,err: ", err.Error())
						}

						//TODO 布隆过滤器介质更换
						r.bloom.Add(formatId)
					}
				}

				marshal, _ := jsoniter.Marshal(value)
				err = repoMessagePusher.KPush(ctx, strconv.Itoa(int(time.Now().Unix())), string(marshal))
				if err != nil {
					log.GetTextLogger().Error("error pushing msg,err: ", err.Error())
				}
			}
		}()
	}
	wg.Wait()
}

func NewRepoCrawl(b bloom.Client, c *bootstrap.KafkaConf, env *bootstrap.Env) domain.RepoCrawl {
	return &repoCrawl{bloom: b, kafkaConf: c, env: env}
}
