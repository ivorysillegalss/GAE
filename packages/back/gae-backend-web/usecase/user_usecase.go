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

func (u *userUsecase) QueryUserNetwork(username string) *domain.UserRelation {
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

func newUserRelation(resp *pb.RelationResponse) *domain.UserRelation {
	member := resp.GetNetMember()
	var nms []*domain.NetMember
	for _, netMember := range member {
		nm := &domain.NetMember{
			Username:  netMember.Username,
			AvatarUrl: netMember.AvatarUrl,
			Evidence:  netMember.AvatarUrl,
		}
		nms = append(nms, nm)
	}

	return &domain.UserRelation{
		Username:  resp.GetUsername(),
		HtmlUrl:   resp.GetHtmlUrl(),
		NetMember: &nms,
	}
}
