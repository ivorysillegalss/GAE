package usecase

import "gae-backend-web/domain"

type rankUsecase struct {
	rankRepository domain.RankRepository
}

func (r *rankUsecase) GetHotRank(page int, phase string) *[]*domain.RankUser {
	return r.rankRepository.CacheGetHotRank(page, phase)
}

func (r *rankUsecase) GetHotRankPhase() *[]string {
	return r.rankRepository.CacheGetRankPhase()
}

func (r *rankUsecase) SearchUserRank(username string) *[]*domain.RankUser {
	//TODO redis打缓存
	return r.rankRepository.SearchUserRank(username)
}

func NewRankUsecase(rr domain.RankRepository) domain.RankUsecase {
	return &rankUsecase{rankRepository: rr}
}
