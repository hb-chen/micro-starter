package client

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	user "github.com/hb-go/micro/auth/srv/proto/user"

	"golang.org/x/net/context"
)

type userKey struct{}

// FromContext retrieves the client from the Context
func UserFromContext(ctx context.Context) (user.UserService, bool) {
	c, ok := ctx.Value(userKey{}).(user.UserService)
	return c, ok
}

// Client returns a wrapper for the PostClient
func UserWrapper(service micro.Service) server.HandlerWrapper {
	client := user.NewUserService("go.micro.srv.auth", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, userKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
