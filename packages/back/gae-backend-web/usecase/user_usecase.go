package usecase

import "gae-backend-web/domain"

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUsecase(ur domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepository: ur}
}

func (u *userUsecase) QueryUserSpecificInfo(username string) *[]*domain.Contributor {
	//TODO 打缓存
	return u.userRepository.QueryUser(username)
}
