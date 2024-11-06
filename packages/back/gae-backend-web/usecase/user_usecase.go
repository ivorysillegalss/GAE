package usecase

import (
	"context"
	"gae-backend-web/bootstrap"
	"gae-backend-web/domain"
	"gae-backend-web/infrastructure/log"
	pb "gae-backend-web/proto/relation/.proto"
	"time"
)

type userUsecase struct {
	userRepository domain.UserRepository
	rpcEngine      *bootstrap.RpcEngine
}

func NewUsecase(ur domain.UserRepository, re *bootstrap.RpcEngine) domain.UserUsecase {
	return &userUsecase{userRepository: ur, rpcEngine: re}
}

func (u *userUsecase) QueryUserSpecificInfo(username string) *[]*domain.Contributor {
	//TODO 打缓存
	return u.userRepository.QueryUser(username)
}

func (u *userUsecase) QueryByNation(nation string, score int) *[]*domain.RankUser {
	return u.userRepository.SearchUserByNation(nation, score)
}

func (u *userUsecase) QueryByTech(tech string, score int) *[]*domain.RankUser {
	return u.userRepository.SearchUserByTech(tech, score)
}

func (u *userUsecase) QueryByGrade(level string, score int) *[]*domain.RankUser {
	return u.userRepository.SearchUserByLevel(level, score)
}

func (u *userUsecase) QueryUserNetwork(username string) *domain.RelationInfo {
	client := *u.rpcEngine.GrpcClient
	serviceClient := pb.NewRelationNetServiceClient(client.GetConn())
	ctxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	p := &pb.RelationRequest{Username: username}
	relationResp, err := serviceClient.QueryRelationNet(ctxt, p)
	if err != nil {
		log.GetTextLogger().Error("relation info query error, username: " + username)
		return nil
	}
	userRelation := newUserRelation(username, relationResp)
	return userRelation
}

func newUserRelation(username string, resp *pb.RelationResponse) *domain.RelationInfo {
	var rm []*domain.RelationMember
	info := resp.GetNetInfo()
	for _, netInfo := range info {
		var m domain.RelationMember
		m.Weight = int(netInfo.GetRelationWeight())
		m.AvatarUrl = netInfo.GetAvatarUrl()
		m.Username = netInfo.GetUsername()
		rm = append(rm, &m)
	}
	return &domain.RelationInfo{
		Username:       username,
		RelationMember: &rm,
	}
}
