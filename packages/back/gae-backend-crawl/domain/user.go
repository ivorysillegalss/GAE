package domain

import (
	"gae-backend-crawl/internal/checkutil"
	"github.com/google/go-github/github"
)

type Contributor struct {
	Login          string
	Id             int64
	AvatarURL      string
	URL            string
	Contributions  int
	HTMLURL        string
	GravatarID     string
	Name           string
	Company        string
	Blog           string
	Location       string
	Email          string
	Bio            string
	PublicRepos    int
	PublicGists    int
	FollowerCount  int
	FollowingCount int
	CreatedAt      int64
	UpdatedAt      int64
	Type           string
	Followers      []*github.User
	Followings     []*github.User
}

func NewContributorValue(g *github.Contributor, v *github.User) *Contributor {
	return &Contributor{
		Name:      checkutil.CheckString(v.Name),
		AvatarURL: checkutil.CheckString(g.AvatarURL),
		Company:   checkutil.CheckString(v.Company),
		Blog:      checkutil.CheckString(v.Blog),
		Location:  checkutil.CheckString(v.Location),
		Email:     checkutil.CheckString(v.Email),
		Bio:       checkutil.CheckString(v.Bio),
		//以上字段有可能为空

		Login:          *g.Login,
		Id:             *g.ID,
		URL:            *g.URL,
		Contributions:  *g.Contributions,
		HTMLURL:        *v.HTMLURL,
		GravatarID:     *v.GravatarID,
		PublicRepos:    *v.PublicRepos,
		PublicGists:    *v.PublicGists,
		FollowerCount:  *v.Followers,
		FollowingCount: *v.Following,
		CreatedAt:      v.CreatedAt.Unix(),
		UpdatedAt:      v.UpdatedAt.Unix(),
		Type:           *v.Type,
	}
}
