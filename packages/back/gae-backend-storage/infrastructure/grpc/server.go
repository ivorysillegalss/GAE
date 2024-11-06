package grpc

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server interface {
	Get() (net.Listener, *grpc.Server)
}

type grpcServer struct {
	GrpcServer *grpc.Server
	Lis        net.Listener
}

func (g *grpcServer) Get() (net.Listener, *grpc.Server) {
	return g.Lis, g.GrpcServer
}

func NewGrpcServer(grpcServerNetWork string, grpcServerAddress string) Server {
	lis, err := net.Listen(grpcServerNetWork, ":"+grpcServerAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	g := grpc.NewServer()
	return &grpcServer{GrpcServer: g, Lis: lis}
}
