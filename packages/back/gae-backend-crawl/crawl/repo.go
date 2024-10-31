package crawl

import (
	"context"
	"gae-backend-crawl/domain"
	"gae-backend-crawl/infrastructure/bloom"
	"gae-backend-crawl/infrastructure/log"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"strconv"
	"sync/atomic"
)

type RepoCrawl struct {
	bloom  bloom.Client
	repoId int64
}

func getRepoId(r *RepoCrawl) int64 {
	return atomic.AddInt64(&r.repoId, 1)
}

func (r *RepoCrawl) DoCrawl() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	repoId := getRepoId(r)
	repo, _, err := client.Repositories.GetByID(ctx, repoId)
	if err != nil {
		log.GetTextLogger().Error("Error fetching repository by Id: %v", err)
	}

	//TODO 贡献者过多时分页多次请求
	opts := &github.ListContributorsOptions{ListOptions: github.ListOptions{PerPage: 100}}
	contributors, _, err1 := client.Repositories.ListContributors(ctx, *repo.Owner.Name, *repo.Name, opts)

	value := domain.NewRepositoryValue(repo)
	for _, contributor := range contributors {

		//1创建新的仓库结构体，将贡献者的列表加到结构体当中
		formatId := strconv.FormatInt(*contributor.ID, 10)
		id := *value.ContributorsId
		id = append(id, formatId)

		//2 查布隆 无记录则加记录 TODO 并生产此用户相关信息
		isExist := r.bloom.Check(formatId)
		if !isExist {
			//TODO
			r.bloom.Add(formatId)
		}
	}
	//TODO 仓库信息生产

	if err1 != nil {
		log.GetTextLogger().Error("Error fetching contributors by Id: %v", err)
	}
}

func NewRepoCrawl(b bloom.Client) *RepoCrawl {
	return &RepoCrawl{bloom: b}
}
