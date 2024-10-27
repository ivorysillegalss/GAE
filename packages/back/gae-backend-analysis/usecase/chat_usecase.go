package usecase

import (
	"context"
	_ "embed"
	"gae-backend-analysis/api/middleware/taskchain"
	"gae-backend-analysis/bootstrap"
	"gae-backend-analysis/constant/cache"
	"gae-backend-analysis/constant/common"
	task2 "gae-backend-analysis/constant/task"
	"gae-backend-analysis/domain"
	"gae-backend-analysis/infrastructure/log"
	"gae-backend-analysis/internal/tokenutil"
	"gae-backend-analysis/task"
)

//go:embed lua/increment.lua
var incrementLuaScript string

type chatUseCase struct {
	env                  *bootstrap.Env
	chatRepository       domain.ChatRepository
	generationRepository domain.GenerationRepository
	chatTask             task.AskTask
	tokenUtil            *tokenutil.TokenUtil
	convertTask          task.ConvertTask
}

func NewChatUseCase(e *bootstrap.Env, c domain.ChatRepository, ct task.AskTask, util *tokenutil.TokenUtil, cvt task.ConvertTask, gr domain.GenerationRepository) domain.ChatUseCase {
	chat := &chatUseCase{chatRepository: c, env: e, chatTask: ct, tokenUtil: util, convertTask: cvt, generationRepository: gr}
	return chat
}

func (c *chatUseCase) InitChat(ctx context.Context, token string, botId int) int {
	//ctx, cancel := context.WithTimeout(ctx, time.Duration(c.env.ContextTimeout))
	//defer cancel()

	script := incrementLuaScript
	log.GetJsonLogger().Info("load new chat lua script")

	chatId, err := c.chatRepository.CacheLuaInsertNewChatId(ctx, script, cache.NewestChatIdKey)
	if err != nil {
		log.GetTextLogger().Error(err.Error())
		return common.FalseInt
	}

	id, err := c.tokenUtil.DecodeToId(token)
	if err != nil {
		log.GetTextLogger().Error(err.Error())
		return common.FalseInt
	}

	// 同样提供依赖mq or not
	go c.chatRepository.DbInsertNewChat(ctx, id, botId)

	return chatId
}

func (c *chatUseCase) ContextChat(ctx context.Context, token string, botId int, chatId int, askMessage string, adjustment bool) (isSuccess bool, message domain.ParsedResponse, code int) {
	chatTask := c.chatTask

	userId, err := c.tokenUtil.DecodeToId(token)
	if err != nil {
		return false, &domain.OpenAIParsedResponse{GenerateText: common.ZeroString}, common.FalseInt
	}

	//我他妈太优雅了
	taskContext := chatTask.InitContextData(userId, botId, chatId, askMessage, task2.ExecuteChatAskType, task2.ExecuteChatAskCode, task2.ChatAskExecutorId, adjustment)
	factory := taskchain.NewTaskContextFactory()
	factory.TaskContext = taskContext
	factory.Puts(chatTask.PreCheckDataTask, chatTask.GetHistoryTask, chatTask.GetBotTask,
		chatTask.AssembleReqTask, chatTask.CallApiTask, chatTask.ParseRespTask, chatTask.StorageTask)
	factory.ExecuteChain()

	//按理来说 上面的taskContext == factory.TaskContext 但是下面再赋值一下比较稳妥一点
	taskContext = factory.TaskContext
	if taskContext.Exception {
		return false, &domain.OpenAIParsedResponse{GenerateText: taskContext.TaskContextResponse.Message}, taskContext.TaskContextResponse.Code
	}
	data := taskContext.TaskContextData.(*domain.AskContextData)
	parsedResponse := data.ParsedResponse

	response := parsedResponse.(*domain.OpenAIParsedResponse)
	return true, response, task2.SuccessCode
}

func (c *chatUseCase) InitMainPage(ctx context.Context, token string, botId int) (titles []*domain.TitleData, err error) {
	userId, err := c.tokenUtil.DecodeToId(token)
	if err != nil {
		return nil, err
	}
	titleStr, err := c.chatRepository.CacheGetTitles(ctx, userId, botId)
	return titleStr, nil
}

func (c *chatUseCase) InputUpdateTitle(ctx context.Context, title string, token string, chatId int, botId int) bool {
	panic("error")
}
