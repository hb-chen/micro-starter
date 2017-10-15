package handler

import (
token "github.com/hb-go/micro/auth/srv/proto/token"
"golang.org/x/net/context"
)

type Token struct{}

func (t *Token) Generate(ctx context.Context, req *token.ReqKey, rsp *token.Rsp) error {
	rsp.Token = "token:"+req.Key
	rsp.Verified = true

	return nil
}

func (t *Token) Verify(ctx context.Context, req *token.ReqToken, rsp *token.Rsp) error {
	rsp.Token = req.Token
	rsp.Verified = true

	return nil
}

