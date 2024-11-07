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
	HTMLURL        string
	GravatarID     string
	Name           string
	Company        string
	Blog           string
	Location       string
	Email          string
	Bio            string
	PublicRepos    int32
	PublicGists    int32
	FollowerCount  int32
	FollowingCount int32
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
		HTMLURL:        *v.HTMLURL,
		GravatarID:     *v.GravatarID,
		PublicRepos:    int32(*v.PublicRepos),
		PublicGists:    int32(*v.PublicGists),
		FollowerCount:  int32(*v.Followers),
		FollowingCount: int32(*v.Following),
		CreatedAt:      v.CreatedAt.Unix(),
		UpdatedAt:      v.UpdatedAt.Unix(),
		Type:           *v.Type,
	}
}
