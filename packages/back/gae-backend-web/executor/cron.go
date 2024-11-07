package executor

import (
	"gae-backend-web/domain"
)

type CronExecutor struct {
	TalentCron domain.TalentCron
}

// SetupCron 启动定时任务
func (d *CronExecutor) SetupCron() {

}

func NewCronExecutor() *CronExecutor {
	return &CronExecutor{}
}
