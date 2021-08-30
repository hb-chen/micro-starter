package handler

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/test"
	"github.com/micro/micro/v3/service/logger"

	pb "github.com/hb-chen/micro/console/srv/proto/user"
)

type User struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) Login(ctx context.Context, req *pb.LoginRequest, rsp *pb.LoginResponse) error {
	logger.Info("Received Example.Call request")

	claims := jwt.StandardClaims{
		Id:        req.Username,
		ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
	}

	privateKey := test.LoadRSAPrivateKeyFromDisk("./conf/auth_key")
	tokenString := test.MakeSampleToken(claims, privateKey)

	rsp.Token = tokenString

	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) Logout(ctx context.Context, req *pb.Request, rsp *pb.LogoutResponse) error {
	logger.Info("Received Example.Call request")
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) Info(ctx context.Context, req *pb.Request, rsp *pb.InfoResponse) error {
	logger.Info("Received Example.Call request")
	rsp.Name = "Hobo"
	rsp.Avatar = "https://avatars3.githubusercontent.com/u/730866?s=460&v=4"
	return nil
}
