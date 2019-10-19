package handler

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/test"
	"github.com/micro/go-micro/util/log"

	example "github.com/hb-go/micro/console/srv/proto/user"
)

type User struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) Login(ctx context.Context, req *example.LoginRequest, rsp *example.LoginResponse) error {
	log.Log("Received Example.Call request")

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
func (e *User) Logout(ctx context.Context, req *example.Request, rsp *example.LogoutResponse) error {
	log.Log("Received Example.Call request")
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) Info(ctx context.Context, req *example.Request, rsp *example.InfoResponse) error {
	log.Log("Received Example.Call request")
	rsp.Name = "Hobo"
	rsp.Avatar = "https://avatars3.githubusercontent.com/u/730866?s=460&v=4"
	return nil
}
