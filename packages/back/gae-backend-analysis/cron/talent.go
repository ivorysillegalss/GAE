package cron

import (
	"gae-backend-analysis/bootstrap"
	"gae-backend-analysis/constant/sys"
	"gae-backend-analysis/domain"
	"gae-backend-analysis/handler"
	"gae-backend-analysis/infrastructure/log"
	jsoniter "github.com/json-iterator/go"
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
		offset, shardValue, err := t.talentRepository.GetAndUpdateCleansingDataShardOffset()
		if err != nil {
			log.GetTextLogger().Error("error in get cleansing data shard offset: " + strconv.Itoa(offset))
		}

		//TODO O(n)复杂度，待优化
		var talents []domain.Talent
		for _, data := range shardValue {
			var talent domain.Talent
			err = jsoniter.Unmarshal([]byte(data), &talent)
			if err != nil {
				log.GetTextLogger().Error("Error unmarshalling JSON: %v", err)
			}
			talents = append(talents, talent)
		}

		h := handler.NewIntermediateHandler()
		var wg sync.WaitGroup
		task := func() {
			defer wg.Done()
			h.WriteData(talents)
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
