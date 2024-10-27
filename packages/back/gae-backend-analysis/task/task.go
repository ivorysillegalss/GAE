package task

import "gae-backend-analysis/api/middleware/taskchain"

// TODO REMOVE同样没有实际意义，此处仅作为责任链壳子
type AskTask interface {
	InitContextData(args ...any) *taskchain.TaskContext
	// PreCheckDataTask 数据的前置检查 & 组装TaskContextData对象
	PreCheckDataTask(tc *taskchain.TaskContext)
	// GetHistoryTask 从DB or Cache获取历史记录
	GetHistoryTask(tc *taskchain.TaskContext)
	// GetBotTask 获取prompt & model
	GetBotTask(tc *taskchain.TaskContext)
	// TODO 微调 TBD
	AdjustmentTask(tc *taskchain.TaskContext)
	// AssembleReqTask 组装rpc请求体
	AssembleReqTask(tc *taskchain.TaskContext)
	// CallApiTask 调用api
	CallApiTask(tc *taskchain.TaskContext)
	// ParseRespTask 转换rpc后响应数据
	ParseRespTask(tc *taskchain.TaskContext)
	// StorageTask 存储相关
	StorageTask(tc *taskchain.TaskContext)
}

// ConvertTask 责任链中节点在不同方法中使用的时候 需要根据需求进行一定定制修改
// 对于变化较小的改动 直接在此定义节点并使用即可 上方title_task那其实也可以这么干
type ConvertTask interface {
	InitStreamStorageTask(args ...any) *taskchain.TaskContext
	StreamArgsTask(tc *taskchain.TaskContext)
	StreamStorageTask(*taskchain.TaskContext)
}
