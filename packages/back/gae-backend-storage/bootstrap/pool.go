package bootstrap

import (
	"gae-backend-storage/constant/sys"
	"gae-backend-storage/infrastructure/pool"
)

func NewPoolFactory() *PoolsFactory {
	p := make(map[int]*pool.Pool, sys.GoRoutinePoolTypesAmount)
	return &PoolsFactory{Pools: p}
}
