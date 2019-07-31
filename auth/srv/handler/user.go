package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	user "github.com/hb-go/micro/auth/srv/proto/user"
)

type User struct{}

func (u *User) GetUser(ctx context.Context, req *user.ReqId, rsp *user.Rsp) error {
	log.Log("Received User.GetUser request")
	rsp.Id = req.Id
	rsp.Nickname = "Hobo"
	return nil
}

func (u *User) GetUserLogin(ctx context.Context, req *user.ReqLogin, rsp *user.Rsp) error {
	log.Log("Received User.GetUserLogin request")
	rsp.Id = 1
	rsp.Nickname = req.Nickname
	return nil
}
