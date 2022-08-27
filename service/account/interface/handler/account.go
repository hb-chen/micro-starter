package handler

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/test"
	"github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"

	"github.com/hb-chen/micro-starter/service/account/conf"
	account "github.com/hb-chen/micro-starter/service/account/proto/account"
	"github.com/hb-chen/micro-starter/service/account/usecase"
)

type Account struct {
	userUseCase usecase.UserUseCase
}

func NewAccountService(userUseCase usecase.UserUseCase) *Account {
	return &Account{
		userUseCase: userUseCase,
	}
}

// Call is a single request handler called via client.Call or the generated client code
func (a *Account) Login(ctx context.Context, req *account.LoginRequest, rsp *account.LoginResponse) error {
	log.Infof("Received Account.Login request")

	user, err := a.userUseCase.LoginUser(req.Username, req.Password)
	if err != nil {
		return err
	} else if user == nil {
		return errors.Forbidden("go.micro.srv.account", "用户名或密码错误: %v, %v", req.Username, req.Password)
	}

	claims := jwt.StandardClaims{
		Id:        strconv.FormatInt(user.Id, 10),
		ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		Subject:   req.Username,
	}

	privateKey := test.LoadRSAPrivateKeyFromDisk(conf.BASE_PATH + "auth_key")
	tokenString := test.MakeSampleToken(claims, jwt.SigningMethodRS256, privateKey)
	rsp.Token = tokenString

	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (*Account) Logout(ctx context.Context, req *account.Request, rsp *account.LogoutResponse) error {
	log.Info("Received Account.Logout request")
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (a *Account) Info(ctx context.Context, req *account.Request, rsp *account.InfoResponse) error {
	log.Info("Received Account.Info request")
	user, err := a.userUseCase.GetUser(req.Id)
	if err != nil {
		return err
	} else if user == nil {
		return errors.NotFound("go.micro.srv.account", "用户不存在: %v", req.Id)
	}

	rsp.Name = fmt.Sprintf("%s-ID:%d", user.Name, req.Id)
	rsp.Avatar = "https://avatars3.githubusercontent.com/u/730866?s=460&v=4"
	return nil
}
