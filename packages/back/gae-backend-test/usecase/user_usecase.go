package usecase

import (
	"gae-backend-test/domain"
	"github.com/bxcodec/faker/v3"
	"math/rand"
)

type userUsecase struct {
}

func NewUsecase() domain.UserUsecase {
	return &userUsecase{}
}

func generateFakeContributor() *[]*domain.Contributor {
	var contributors []*domain.Contributor
	for i := 0; i < 10; i++ {
		var contributor domain.Contributor
		_ = faker.FakeData(&contributor) // 自动生成假数据
		contributors = append(contributors, &contributor)
	}
	return &contributors
}
func (u *userUsecase) QueryUserSpecificInfo(username string) *[]*domain.Contributor {
	return generateFakeContributor()
}

func (u *userUsecase) QueryUserNetwork(username string) *domain.RelationInfo {
	member := &domain.RelationMember{
		Username:  faker.Username(),
		Weight:    666,
		AvatarUrl: faker.URL(),
	}
	var ms []*domain.RelationMember
	members := append(ms, member)
	return &domain.RelationInfo{
		RelationMember: &members,
		Username:       username,
	}
}

func (u *userUsecase) QueryByNation(nation string, score int) *[]*domain.RankUser {
	return generateRankUsers(nation, score)
}

func (u *userUsecase) QueryByTech(tech string, score int) *[]*domain.RankUser {
	return generateRankUsers(tech, score)
}

func (u *userUsecase) QueryByGrade(level string, score int) *[]*domain.RankUser {
	return generateRankUsers(level, score)
}

func generateRankUsers(nation string, score int) *[]*domain.RankUser {
	var rankUsers []*domain.RankUser

	// Generate a fixed number of RankUser entries, e.g., 10
	for i := 0; i < 10; i++ {
		rankUsers = append(rankUsers, generateFakeRankUser(nation, score))
	}

	return &rankUsers
}

// generateFakeRankUser generates a single RankUser with random data
func generateFakeRankUser(nation string, score int) *domain.RankUser {
	return &domain.RankUser{
		Placement:              666,
		Login:                  faker.Username(),
		Id:                     rand.Int63n(1000000), // Random ID for testing
		AvatarURL:              faker.URL(),
		HTMLURL:                faker.URL(),
		Location:               "nation", // Use provided nation for testing
		Company:                "company",
		Score:                  123,
		Level:                  generateRandomLevel(), // Generate a random level
		LocationClassification: 0.8,
		Tech:                   faker.Word(),
	}
}

// generateRandomLevel returns a random level for testing
func generateRandomLevel() string {
	levels := []string{"A", "B", "C", "D"}
	return levels[rand.Intn(len(levels))]
}
