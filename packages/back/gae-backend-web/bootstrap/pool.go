package bootstrap

import (
	"gae-backend-web/constant/sys"
	"gae-backend-web/infrastructure/pool"
)

func NewPoolFactory() *PoolsFactory {
	p := make(map[int]*pool.Pool, sys.GoRoutinePoolTypesAmount)
	return &PoolsFactory{Pools: p}
}
