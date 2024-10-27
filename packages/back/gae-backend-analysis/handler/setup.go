package handler

import (
	"gae-backend-analysis/bootstrap"
	task2 "gae-backend-analysis/constant/task"
	"gae-backend-analysis/domain"
	"gae-backend-analysis/infrastructure/log"
)

func NewLanguageModelExecutor(env *bootstrap.Env, channels *bootstrap.Channels, executorId int) domain.LanguageModelExecutor {
	var executor domain.LanguageModelExecutor
	var executorType string
	switch executorId {
	case task2.ChatAskExecutorId:
		executor = &OpenaiChatModelExecutor{env: env, res: channels}
		executorType = task2.ExecuteChatAskType
	case task2.ChatTitleAskExecutorId:
		executor = &OpenaiChatModelExecutor{env: env, res: channels}
		executorType = task2.ExecuteTitleAskType
	default:
		log.GetTextLogger().Fatal("illegal llm executor id")
	}

	log.GetJsonLogger().WithFields("choose executor", true, "executor type", executorType)
	return executor
}
