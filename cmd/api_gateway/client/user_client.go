package client

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "sports-go/shared/pb/user"
)

// UserClient 封裝用戶服務的 gRPC client
type UserClient struct {
	conn   *grpc.ClientConn
	client pb.UserServiceClient
}

// NewUserClient 創建新的用戶服務 client
func NewUserClient(serverAddr string) (*UserClient, error) {
	// 建立 gRPC 連接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to user service: %v", err)
	}

	return &UserClient{
		conn:   conn,
		client: pb.NewUserServiceClient(conn),
	}, nil
}

// Close 關閉連接
func (c *UserClient) Close() error {
	return c.conn.Close()
}

// GetUser 獲取用戶信息
func (c *UserClient) GetUser(ctx context.Context, userID string) (*pb.GetUserResp, error) {
	req := &pb.GetUserReq{
		UserId: userID,
	}
	return c.client.GetUser(ctx, req)
}

// Login 用戶登錄
func (c *UserClient) Login(ctx context.Context, userID, token string) (*pb.LoginResp, error) {
	req := &pb.LoginReq{
		UserId: userID,
		Token:  token,
	}
	return c.client.Login(ctx, req)
}
