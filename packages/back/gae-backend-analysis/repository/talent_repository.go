package repository

import (
	"context"
	"gae-backend-analysis/constant/cache"
	"gae-backend-analysis/constant/common"
	"gae-backend-analysis/constant/mq"
	"gae-backend-analysis/domain"
	"gae-backend-analysis/infrastructure/redis"
	"github.com/thoas/go-funk"
	"strconv"
)

type talentRepository struct {
	rcl redis.Client
}

func NewTalentRepository(rcl redis.Client) domain.TalentRepository {
	return &talentRepository{rcl: rcl}
}

// GetAndUpdateCleansingDataShardOffset 已知存在部分bug
func (t *talentRepository) GetAndUpdateCleansingDataShardOffset(businessId int) (bool, int, []string, error) {
	ctx := context.Background()
	switch businessId {
	case mq.UnCleansingUserId:
		pop, err := t.rcl.RPop(ctx, cache.TemporaryUnRankCleansingContributorReadyShard)
		a, _ := strconv.Atoi(pop)
		if t.rcl.IsEmpty(err) {
			return true, -1, nil, err
		}
		if err != nil {
			return false, a, nil, err
		}
		members, err1 := t.rcl.SMembers(ctx, cache.TemporaryUnRankCleansingContributorData+common.Infix+cache.TemporaryUnRankCleansingContributorShard+common.Infix+pop)
		return false, a, members, err1

	case mq.UnCleansingRepoId:
		pop, err := t.rcl.RPop(ctx, cache.TemporaryUnRankCleansingRepoReadyShard)
		a, _ := strconv.Atoi(pop)
		if t.rcl.IsEmpty(err) {
			return true, -1, nil, err
		}
		if err != nil {
			return false, a, nil, err
		}
		members, err1 := t.rcl.SMembers(ctx, cache.TemporaryUnRankCleansingRepoData+common.Infix+cache.TemporaryUnRankCleansingRepoShard+common.Infix+pop)
		return false, a, members, err1
	default:
		panic("error storage")
	}
}

//	TODO 并发
//
// CleansingDataTemporaryStorageCache 分片存储
func (t *talentRepository) CleansingDataTemporaryStorageCache(header string, value string, businessId int) error {
	ctx := context.Background()
	switch businessId {
	case mq.UnCleansingRepoId:
		listIdStr, _ := t.rcl.Get(ctx, cache.TemporaryUnRankCleansingRepoShardList)
		storagePre := cache.TemporaryUnRankCleansingRepoData + common.Infix + cache.TemporaryUnRankCleansingRepoShard + common.Infix
		card := t.rcl.SCard(ctx, storagePre+listIdStr)
		listId, _ := strconv.Atoi(listIdStr)
		if funk.Equal(card, cache.TemporaryUnRankCleansingRepoShardMaxValue) {
			_ = t.rcl.LPush(ctx, cache.TemporaryUnRankCleansingRepoReadyShard, listIdStr)
			listId += 1
			_ = t.rcl.Set(ctx, cache.TemporaryUnRankCleansingRepoShardList, strconv.Itoa(listId))
		}
		return t.rcl.SAdd(ctx, storagePre+strconv.Itoa(listId), value)

	case mq.UnCleansingUserId:
		listIdStr, _ := t.rcl.Get(ctx, cache.TemporaryUnRankCleansingContributorShardList)
		storagePre := cache.TemporaryUnRankCleansingContributorData + common.Infix + cache.TemporaryUnRankCleansingContributorShard + common.Infix
		card := t.rcl.SCard(ctx, storagePre+listIdStr)
		listId, _ := strconv.Atoi(listIdStr)
		if funk.Equal(card, cache.TemporaryUnRankCleansingContributorShardMaxValue) {
			_ = t.rcl.LPush(ctx, cache.TemporaryUnRankCleansingContributorReadyShard, listIdStr)
			listId += 1
			_ = t.rcl.Set(ctx, cache.TemporaryUnRankCleansingRepoShardList, strconv.Itoa(listId))
		}
		return t.rcl.SAdd(ctx, storagePre+strconv.Itoa(listId), value)
	default:
		panic("illegal storage")
	}

}
