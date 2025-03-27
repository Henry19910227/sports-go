package api

import (
	"context"
	"internal/service"
	"sports-go/shared/pb"
)

type User struct {
	pb.UnimplementedUserServiceServer
}

func (u *User) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserResp, error) {
	service.Test()
	res := &pb.GetUserResp{
		UserId: req.UserId,
		Name:   "Henry",
		Email:  "email@email.com",
	}
	return res, nil
}

func (u *User) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	service.Test()
	res := &pb.LoginResp{
		Success: true,
		Message: req.UserId + " login success",
	}
	return res, nil
}
