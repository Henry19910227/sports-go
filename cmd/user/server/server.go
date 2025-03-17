package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sports-go/cmd/user/config"
	"sports-go/cmd/user/proto"
)

// Server 定義 gRPC 伺服器
type Server struct {
	grpcServer *grpc.Server
	listener   net.Listener
}

// NewServer 建立 gRPC 伺服器
func NewServer(cfg *config.Config, user *proto.User) *Server {
	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", cfg.GRPCPort, err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, user)
	reflection.Register(grpcServer) // 啟用 gRPC 反射（方便 gRPC UI 使用）

	return &Server{
		grpcServer: grpcServer,
		listener:   lis,
	}
}

// Start 啟動 gRPC 服務
func (s *Server) Start() error {
	return s.grpcServer.Serve(s.listener)
}

// Stop 優雅關閉 gRPC 服務
func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}
