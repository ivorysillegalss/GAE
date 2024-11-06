package bootstrap

import "gae-backend-storage/infrastructure/grpc"

func NewGrpc(env *Env) grpc.Client {
	return grpc.NewGrpcClient(env.GrpcUrl)
}

func NewGrpcServer(env *Env) grpc.Server {
	return grpc.NewGrpcServer(env.GrpcServerNetwork, env.GrpcServerAddress)
}
