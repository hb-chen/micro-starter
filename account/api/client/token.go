package client

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	token "github.com/hb-go/micro/auth/srv/proto/token"

	"golang.org/x/net/context"
)

type tokenKey struct{}

// FromContext retrieves the client from the Context
func TokenFromContext(ctx context.Context) (token.TokenService, bool) {
	c, ok := ctx.Value(tokenKey{}).(token.TokenService)
	return c, ok
}

// Client returns a wrapper for the TokenClient
func TokenWrapper(service micro.Service) server.HandlerWrapper {
	client := token.NewTokenService("go.micro.srv.auth", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, tokenKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
