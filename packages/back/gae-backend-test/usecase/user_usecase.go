package usecase

import (
	"gae-backend-test/domain"
	"github.com/bxcodec/faker/v3"
	"math/rand"
	"time"
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

func (u *userUsecase) QueryUserNetwork(username string) *[]*domain.NestedNetInfo {
	rand.Seed(time.Now().UnixNano()) // Seed for randomness
	var nns []*domain.NestedNetInfo
	nns = append(nns, GenerateFakeNestedNetInfo())
	return &nns
}

// GenerateFakeNationInfo generates random NationInfo data
func GenerateFakeNationInfo() *domain.NationInfo {
	return &domain.NationInfo{
		Nation:       faker.Word(),
		NationWeight: rand.Int63n(100),
	}
}

// GenerateFakeNestedNetInfo generates a nested structure with depth 2 and length 1
func GenerateFakeNestedNetInfo() *domain.NestedNetInfo {
	// Generate main node (center node)
	mainNode := &domain.NestedNetInfo{
		Relation:   faker.Word(),
		Username:   faker.Username(),
		NationInfo: GenerateFakeNationInfo(),
	}

	// Generate only one child node to fulfill the length constraint
	childNode1 := &domain.NestedNetInfo{
		Relation:   faker.Word(),
		Username:   faker.Username(),
		NationInfo: GenerateFakeNationInfo(),
	}
	childNode2 := &domain.NestedNetInfo{
		Relation:   faker.Word(),
		Username:   faker.Username(),
		NationInfo: GenerateFakeNationInfo(),
	}

	// Assign the single child to the main node's relations
	mainNode.Relations = []*domain.NestedNetInfo{childNode1, childNode2}

	return mainNode
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
		Login:     faker.Username(),
		Id:        rand.Int63n(1000000), // Random ID for testing
		AvatarURL: faker.URL(),
		HTMLURL:   faker.URL(),
		Location:  "nation", // Use provided nation for testing
		Company:   "company",
		Score:     123,
		Level:     generateRandomLevel(), // Generate a random level
	}
}

// generateRandomLevel returns a random level for testing
func generateRandomLevel() string {
	levels := []string{"A", "B", "C", "D"}
	return levels[rand.Intn(len(levels))]
}
