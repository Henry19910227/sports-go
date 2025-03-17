package proto

import (
	"context"
	"internal/service"
)

type User struct {
	UnimplementedUserServiceServer
}

func (u *User) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	service.Test()
	res := &GetUserResponse{
		UserId: req.UserId,
		Name:   "",
		Email:  "",
	}
	return res, nil
}
