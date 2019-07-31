package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	token "github.com/hb-go/micro/auth/srv/proto/token"
)

type Token struct{}

func (t *Token) Generate(ctx context.Context, req *token.ReqKey, rsp *token.Rsp) error {
	log.Logf("Token.Generate with key:%v", req.Key)
	rsp.Token = "token:" + req.Key
	rsp.Verified = true

	return nil
}

func (t *Token) Verify(ctx context.Context, req *token.ReqToken, rsp *token.Rsp) error {
	log.Logf("Token.Verify with token:%v", req.Token)
	rsp.Token = req.Token
	rsp.Verified = true

	return nil
}
