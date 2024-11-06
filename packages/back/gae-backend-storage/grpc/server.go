package grpc

import (
	"context"
	pb "gae-backend-storage/proto/storage/.proto" // 替换为你生成的路径
	"log"
)

// Server 实现了 StorageService 的接口
type Server struct {
	pb.UnimplementedStorageServiceServer
}

func (s *Server) UpdateData(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	log.Printf("Received data: %s", req.Data)
	// 模拟存储更新操作
	return &pb.UpdateResponse{Status: "Update successful"}, nil
}
