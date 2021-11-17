package handler

import (
	"context"
	"github.com/micro/micro/v3/service/logger"

	pb "github.com/hb-chen/micro-starter/service/account/proto/account"
)

type Account struct{}

func (a *Account) Login(ctx context.Context, req *pb.LoginRequest, rsp *pb.LoginResponse) error {
	logger.Info("Received Account.Login request")

	rsp.Token = "token"

	return nil
}

func (a *Account) Logout(ctx context.Context, req *pb.Request, rsp *pb.LogoutResponse) error {
	logger.Info("Received Account.Logout request")
	return nil
}

func (a *Account) Info(ctx context.Context, req *pb.Request, rsp *pb.InfoResponse) error {
	logger.Info("Received Account.Info request")
	rsp.Name = "Hobo"
	rsp.Avatar = "https://avatars3.githubusercontent.com/u/730866?s=460&v=4"
	return nil
}
