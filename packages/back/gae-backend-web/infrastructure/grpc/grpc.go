package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type Client interface {
	GetConn() *grpc.ClientConn
}

type grpcClient struct {
	*grpc.ClientConn
}

func (g *grpcClient) GetConn() *grpc.ClientConn {
	return g.ClientConn
}

func NewGrpcClient(grpcUrl string) Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	conn, err := grpc.DialContext(ctx, grpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	return &grpcClient{ClientConn: conn}
}
