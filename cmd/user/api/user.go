package api

import (
	"context"
	"internal/service"
	"sports-go/shared/pb/user"
)

type User struct {
	user.UnimplementedUserServiceServer
}

func (u *User) GetUser(ctx context.Context, req *user.GetUserReq) (*user.GetUserResp, error) {
	service.Test()
	res := &user.GetUserResp{
		UserId: req.UserId,
		Name:   "Henry",
		Email:  "email@email.com",
	}
	return res, nil
}

func (u *User) Login(ctx context.Context, req *user.LoginReq) (*user.LoginResp, error) {
	service.Test()
	res := &user.LoginResp{
		Success: true,
		Message: req.UserId + " login success",
	}
	return res, nil
}
