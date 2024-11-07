package grpc

import (
	"context"
	"gae-backend-storage/bootstrap"
	pb "gae-backend-storage/proto/storage/.proto"
	"log"
)

func RegisterGrpcServer(rpc *bootstrap.RpcEngine) {

	//serverPtr := *rpc.GrpcServer
	//lis, server := serverPtr.Get()
	//
	//pb.RegisterStorageServiceServer(server, &Server{})
	//log.Println("Server is running on port 50051")
	//if err := server.Serve(lis); err != nil {
	//	log.Fatalf("Failed to serve: %v", err)
	//}
}

type Server struct {
	pb.UnimplementedStorageServiceServer
}

func (s *Server) UpdateData(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	log.Printf("Received data: %s", req.Data)
	// 模拟存储更新操作
	return &pb.UpdateResponse{Status: "Update successful"}, nil
}
