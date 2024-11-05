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
	return &userUsecase{userRepository: ur}
}

func (u *userUsecase) QueryUserSpecificInfo(username string) *[]*domain.Contributor {
	//TODO 打缓存
	return u.userRepository.QueryUser(username)
}

func (u *userUsecase) QueryUserNetwork(username string) *[]*domain.NestedNetInfo {
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
	userRelation := newUserRelation(relationResp)
	return userRelation
}

func (u *userUsecase) QueryByNation(username string) *[]*domain.Contributor {
	panic("s")
}

func (u *userUsecase) QueryByTech(username string) *[]*domain.Contributor {
	//TODO implement me
	panic("implement me")
}

func newUserRelation(resp *pb.RelationResponse) *[]*domain.NestedNetInfo {
	// 使用map分组NetInfo
	groupedData := make(map[string][]domain.NetInfo)

	// 遍历响应中的每个 NetInfo 并按 Username1 进行分组
	for _, info := range resp.GetNetInfo() {
		ni := domain.NetInfo{
			Relation:  info.GetRelation(),
			Username1: info.GetUsername1(),
			Username2: info.GetUsername2(),
		}
		groupedData[ni.Username1] = append(groupedData[ni.Username1], ni)
	}

	// 将 map 转换为 NestedNetInfo 的列表
	var nestedRelations []*domain.NestedNetInfo
	for username, relations := range groupedData {
		nestedRelations = append(nestedRelations, &domain.NestedNetInfo{
			Username1: username,
			Relations: relations,
		})
	}

	return &nestedRelations
}
