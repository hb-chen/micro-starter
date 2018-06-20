package client

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	comment "github.com/hb-go/micro/post/srv/proto/comment"

	"golang.org/x/net/context"
)

type commentKey struct{}

// FromContext retrieves the client from the Context
func CommentFromContext(ctx context.Context) (comment.CommentService, bool) {
	c, ok := ctx.Value(commentKey{}).(comment.CommentService)
	return c, ok
}

// Client returns a wrapper for the CommentClient
func CommentWrapper(service micro.Service) server.HandlerWrapper {
	client := comment.NewCommentService("go.micro.srv.post", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, commentKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
