package domain

import "github.com/google/go-github/github"

type Contributor struct {
	Login         string
	Id            int64
	AvatarURL     *string
	URL           string
	Contributions int
}

func NewContributorValue(g *github.Contributor) *Contributor {
	return &Contributor{
		Login:         *g.Login,
		Id:            *g.ID,
		AvatarURL:     g.AvatarURL,
		URL:           *g.URL,
		Contributions: *g.Contributions,
	}
}
