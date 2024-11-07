package bootstrap

import "gae-backend-web/infrastructure/grpc"

func NewGrpc(env *Env) grpc.Client {
	return grpc.NewGrpcClient(env.GrpcUrl)
}
