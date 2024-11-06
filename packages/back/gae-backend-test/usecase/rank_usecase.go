package usecase

import "gae-backend-test/domain"

type rankUsecase struct {
	rankRepository domain.RankRepository
}

func (r *rankUsecase) GetHotRank(page int, phase string) *[]*domain.RankUser {
	return generateRankUsers(phase, page)
}

func (r *rankUsecase) GetHotRankPhase() *[]string {
	return &[]string{"1", "2", "3"}
}

func (r *rankUsecase) SearchUserRank(username string) *[]*domain.RankUser {
	return generateRankUsers(username, 1)
}

func NewRankUsecase() domain.RankUsecase {
	return &rankUsecase{}
}
