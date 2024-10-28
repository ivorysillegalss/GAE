package cron

import (
	"gae-backend-analysis/bootstrap"
	"gae-backend-analysis/constant/sys"
	"gae-backend-analysis/domain"
	"gae-backend-analysis/infrastructure/log"
	"strconv"
	"sync"
	"time"
)

type talentCron struct {
	talentRepository domain.TalentRepository
	poolFactory      *bootstrap.PoolsFactory
}

func NewTalentCron(t domain.TalentRepository, p *bootstrap.PoolsFactory) domain.TalentCron {
	return &talentCron{talentRepository: t, poolFactory: p}
}

func (t *talentCron) AnalyseTalent() {
	for {
		offset, err := t.talentRepository.GetAndUpdateCleansingDataShardOffset()
		if err != nil {
			log.GetTextLogger().Error("error in get cleansing data shard offset: " + strconv.Itoa(offset))
		}
		//TODO TBD
		var wg sync.WaitGroup
		task := func() {
			//TODO 此处RPC调用算法API，按需包装执行器
			defer wg.Done()
		}
		config := t.poolFactory.Pools[sys.ExecuteTalentAnalysis]
		wg.Add(1)
		err1 := config.Submit(task)
		if err1 != nil {
			log.GetTextLogger().Error("task pool upload error")
		}
		log.GetTextLogger().Info("success commit analysis task")

		//10s执行一次
		time.Sleep(10 * time.Second)
	}
}
