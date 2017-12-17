package handler

import (
	"encoding/json"

	"github.com/micro/go-log"
	"github.com/micro/go-micro/errors"

	"github.com/hb-go/micro/account/api/client"
	account "github.com/hb-go/micro/account/api/proto/account"
	user "github.com/hb-go/micro/auth/srv/proto/user"
	token "github.com/hb-go/micro/auth/srv/proto/token"
	api "github.com/micro/go-api/proto"
	"golang.org/x/net/context"
	"strings"
)

type Account struct{}

func (a *Account) Login(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Account.Login request")
	log.Logf("req get:%v", req.Get)
	log.Logf("req post:%v", req.Post)
	log.Logf("req body:%v", req.Body)

	reqLogin := &user.ReqLogin{}
	err := json.Unmarshal([]byte(req.Body), reqLogin)
	if err != nil {
		return errors.InternalServerError("go.micro.api.account", "json parse error with:"+err.Error())
	}

	if len(reqLogin.Nickname) == 0 || len(reqLogin.Pwd) == 0 {
		return errors.InternalServerError("go.micro.api.account", "nickname/pwd nil")
	}

	userClient, ok := client.UserFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.account", "user client not found")
	}

	u, err := userClient.GetUserLogin(ctx, reqLogin)
	if err != nil {
		return errors.InternalServerError("go.micro.api.account", "user login err:"+err.Error())
	}

	tokenClient, ok := client.TokenFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.account", "token client not found")
	}

	reqToken := token.ReqKey{Key: u.Nickname}
	t, err := tokenClient.Generate(ctx, &reqToken)
	if err != nil {
		return errors.InternalServerError("go.micro.api.account", "token generate err:"+err.Error())
	}

	response := account.Rsp{}
	response.Id = u.Id
	response.Nickname = u.Nickname
	response.Token = t.Token

	b, _ := json.Marshal(response)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}

func (a *Account) Register(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Account.Register request")

	if strings.Compare("pwd", "pwd") == 0 {
		response := account.Rsp{}
		response.Id = 1
		response.Nickname = "Hobo"
		response.Token = "token"

		b, _ := json.Marshal(response)

		rsp.StatusCode = 200
		rsp.Body = string(b)
		return nil
	} else {
		return errors.InternalServerError("go.micro.api.account", "密码设置不一致，请再次确认")
	}
}
