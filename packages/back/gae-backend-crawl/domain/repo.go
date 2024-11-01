package domain

import (
	"github.com/google/go-github/github"
	"strconv"
)

type RepositoryValue struct {
	RepoId           int64
	OwnerName        string
	OwnerId          string
	Name             string
	CreatedAt        int64
	UpdatedAt        int64
	ForksCount       int
	NetworkCount     int
	OpenIssuesCount  int
	StargazersCount  int
	SubscribersCount int
	WatchersCount    int
	//Contributors     *[]ContributorInfo
	ContributorsId *[]string
}

type ContributorInfo struct {
	ContributorId string
	Contributions int
}

func NewRepositoryValue(repo *github.Repository) *RepositoryValue {
	return &RepositoryValue{
		RepoId:           *repo.ID,
		OwnerName:        *repo.Owner.Name,
		OwnerId:          strconv.FormatInt(*repo.Owner.ID, 10),
		Name:             *repo.Name,
		CreatedAt:        repo.CreatedAt.Time.Unix(),
		UpdatedAt:        repo.UpdatedAt.Time.Unix(),
		ForksCount:       *repo.ForksCount,
		NetworkCount:     *repo.NetworkCount,
		OpenIssuesCount:  *repo.OpenIssuesCount,
		StargazersCount:  *repo.StargazersCount,
		SubscribersCount: *repo.SubscribersCount,
		WatchersCount:    *repo.WatchersCount,
	}
}

type RepoCrawl interface {
	DoCrawl()
}
