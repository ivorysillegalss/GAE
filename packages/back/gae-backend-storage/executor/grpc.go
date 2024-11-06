package executor

import (
	"context"
	"gae-backend-storage/bootstrap"
	"gae-backend-storage/infrastructure/log"
	pb "gae-backend-storage/proto/storage/.proto"
	"time"
)

func queryRanks(e *bootstrap.RpcEngine) bool {
	client := e.GrpcClient
	serviceClient := pb.NewStorageServiceClient(client.GetConn())
	ctxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	p := &pb.UpdateRequest{Data: "active"}
	resp, err := serviceClient.UpdateData(ctxt, p)
	if resp.Status == "success" {
		return true
	}
	if err != nil {
		log.GetTextLogger().Error(err.Error())
	}
	return false
}

type GrpcExecutor struct {
	Rpc          *bootstrap.RpcEngine
	SearchEngine *bootstrap.SearchEngine
}

func (d *GrpcExecutor) SetupGrpc() {
	isSuccess := queryRanks(d.Rpc)
	if isSuccess {
		//TODO
	}
}

func NewGrpcExecutor(e *bootstrap.RpcEngine, es *bootstrap.SearchEngine) *GrpcExecutor {
	return &GrpcExecutor{Rpc: e, SearchEngine: es}
}
