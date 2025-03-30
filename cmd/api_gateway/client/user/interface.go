package user

import (
	"context"
	pb "sports-go/shared/pb/user"
)

type Client interface {
	Close() error
	GetUser(ctx context.Context, userID string) (*pb.GetUserResp, error)
	Login(ctx context.Context, userID int64, token string) (*pb.LoginResp, error)
}
