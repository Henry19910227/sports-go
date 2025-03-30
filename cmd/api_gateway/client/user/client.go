package user

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	pb "sports-go/shared/pb/user"
	"time"
)

type client struct {
	conn   *grpc.ClientConn
	client pb.UserServiceClient
}

func New(address string) (Client, error) {
	// 建立 gRPC 連接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    10 * time.Second, // 每 10 秒探測一次
			Timeout: 5 * time.Second,  // 超時則重新連線
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to user service: %v", err)
	}

	return &client{
		conn:   conn,
		client: pb.NewUserServiceClient(conn),
	}, nil
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) GetUser(ctx context.Context, userID string) (*pb.GetUserResp, error) {
	req := &pb.GetUserReq{
		UserId: userID,
	}
	return c.client.GetUser(ctx, req)
}

func (c *client) Login(ctx context.Context, userID int64, token string) (*pb.LoginResp, error) {
	req := &pb.LoginReq{
		Uid:   userID,
		Token: token,
	}
	return c.client.Login(ctx, req)
}
