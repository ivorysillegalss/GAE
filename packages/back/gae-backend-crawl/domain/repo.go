package domain

import "github.com/google/go-github/github"

type RepositoryValue struct {
	ContributorsId *[]string
}

func NewRepositoryValue(*github.Repository) *RepositoryValue {
	return &RepositoryValue{}
}
