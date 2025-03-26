package api

import (
	"context"
	"internal/service"
)

type User struct {
	UnimplementedUserServiceServer
}

func (u *User) GetUser(ctx context.Context, req *GetUserReq) (*GetUserResp, error) {
	service.Test()
	res := &GetUserResp{
		UserId: req.UserId,
		Name:   "Henry",
		Email:  "email@email.com",
	}
	return res, nil
}

func (u *User) Login(ctx context.Context, req *LoginReq) (*LoginResp, error) {
	service.Test()
	res := &LoginResp{
		Success: true,
		Message: req.UserId + " login success",
	}
	return res, nil
}
