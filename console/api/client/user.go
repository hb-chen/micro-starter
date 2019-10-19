package client

import (
	"context"

	user "github.com/hb-go/micro/console/srv/proto/user"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
)

type exampleKey struct{}

// FromContext retrieves the client from the Context
func UserFromContext(ctx context.Context) (user.UserService, bool) {
	c, ok := ctx.Value(exampleKey{}).(user.UserService)
	return c, ok
}

// Client returns a wrapper for the ExampleClient
func UserWrapper(service micro.Service) server.HandlerWrapper {
	client := user.NewUserService("go.micro.srv.console", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, exampleKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
