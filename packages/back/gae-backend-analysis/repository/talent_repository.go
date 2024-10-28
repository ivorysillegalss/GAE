package repository

import (
	"context"
	"gae-backend-analysis/constant/cache"
	"gae-backend-analysis/constant/common"
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

func (t *talentRepository) GetAndUpdateCleansingDataShardOffset() (int, error) {
	ctx := context.Background()
	pop, err := t.rcl.RPop(ctx, cache.TemporaryUnRankCleansingReadyShard)
	a, _ := strconv.Atoi(pop)
	return a, err
}

// CleansingDataTemporaryStorageCache 分片存储
func (t *talentRepository) CleansingDataTemporaryStorageCache(header string, value string) error {
	ctx := context.Background()
	listIdStr, _ := t.rcl.Get(ctx, cache.TemporaryUnRankCleansingShardList)
	storagePre := cache.TemporaryUnRankCleansingData + common.Infix + cache.TemporaryUnRankCleansingShard + common.Infix
	card := t.rcl.SCard(ctx, storagePre+listIdStr)
	listId, _ := strconv.Atoi(listIdStr)
	if funk.Equal(card, cache.TemporaryUnRankCleansingShardMaxValue) {
		_ = t.rcl.LPush(ctx, cache.TemporaryUnRankCleansingReadyShard, listIdStr)
		listId += 1
	}
	return t.rcl.SAdd(ctx, storagePre+strconv.Itoa(listId), value)
}
