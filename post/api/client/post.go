package client

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	post "github.com/hb-go/micro/post/srv/proto/post"

	"golang.org/x/net/context"
)

type postKey struct{}

// FromContext retrieves the client from the Context
func PostFromContext(ctx context.Context) (post.PostService, bool) {
	c, ok := ctx.Value(postKey{}).(post.PostService)
	return c, ok
}

// Client returns a wrapper for the PostClient
func PostWrapper(service micro.Service) server.HandlerWrapper {
	client := post.NewPostService("go.micro.srv.post", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, postKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
