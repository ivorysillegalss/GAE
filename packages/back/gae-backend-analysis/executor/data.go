package executor

import (
	"context"
	"gae-backend-analysis/constant/cache"
	"gae-backend-analysis/constant/common"
	"gae-backend-analysis/infrastructure/redis"
)

type DataExecutor struct {
	rcl redis.Client
}

func NewDataExecutor(r redis.Client) *DataExecutor {
	return &DataExecutor{rcl: r}
}

var ctx = context.Background()

func initRedisData(d *DataExecutor) {
	_ = d.rcl.Set(ctx, cache.TemporaryUnRankCleansingRepoShardList, common.ZeroInt)
	_ = d.rcl.Set(ctx, cache.TemporaryUnRankCleansingContributorShardList, common.ZeroInt)
}

func (d *DataExecutor) InitData() {
	initRedisData(d)
}
