package bootstrap

import (
	"gae-backend-crawl/constant/sys"
	"gae-backend-crawl/infrastructure/pool"
	"github.com/panjf2000/ants/v2"
)

func NewPoolFactory() *PoolsFactory {
	p := make(map[int]*pool.Pool, sys.GoRoutinePoolTypesAmount)
	defaultPool, _ := ants.NewPool(sys.DefaultPoolGoRoutineAmount) // 1000 可以是您期望的 goroutine 数量
	p[sys.ExecuteRpcGoRoutinePool] = &pool.Pool{Pool: defaultPool}

	return &PoolsFactory{Pools: p}
}
