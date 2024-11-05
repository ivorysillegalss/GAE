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
	return &userRelation
}

func (u *userUsecase) QueryByNation(nation string, score int) *[]*domain.RankUser {
	return u.userRepository.SearchUserByNation(nation, score)
}

func (u *userUsecase) QueryByTech(tech string, score int) *[]*domain.RankUser {
	return u.userRepository.SearchUserByTech(tech, score)
}

func newUserRelation(resp *pb.RelationResponse) []*domain.NestedNetInfo {
	// 使用 map 分组 NetInfo
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

	// 创建最终的嵌套结构体列表
	var nestedRelations []*domain.NestedNetInfo

	// 处理 NationInfo1 和 NationInfo2，获取 Nation 信息
	nationMap := make(map[string]string)
	if resp.GetNationInfo1() != nil {
		nationMap[resp.GetNationInfo1().Nation] = resp.GetNationInfo1().Nation
	}
	if resp.GetNationInfo2() != nil {
		nationMap[resp.GetNationInfo2().Nation] = resp.GetNationInfo2().Nation
	}

	// 将 map 转换为 NestedNetInfo 的列表，并添加 Nation 信息
	for username, relations := range groupedData {
		nestedNetInfo := &domain.NestedNetInfo{
			Username1: username,
			Relations: relations,
		}
		// 如果存在 Nation 信息，赋值给 Nation 字段
		if nation, exists := nationMap[username]; exists {
			nestedNetInfo.Nation = nation
		}
		nestedRelations = append(nestedRelations, nestedNetInfo)
	}

	return nestedRelations
}

func (u *userUsecase) QueryByGrade(level string, score int) *[]*domain.RankUser {
	return u.userRepository.SearchUserByLevel(level, score)
}
