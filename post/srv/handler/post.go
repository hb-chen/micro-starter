package handler

import (
	"github.com/micro/go-log"

	post "github.com/hb-go/micro/post/srv/proto/post"
	"golang.org/x/net/context"
)

type Post struct{}

// Call is a single request handler called via client.GetPost or the generated client code
func (e *Post) GetPost(ctx context.Context, req *post.Req, rsp *post.Rsp) error {
	log.Log("Received Post.GetPost request")
	rsp.Title = "title"
	rsp.Content = "content"
	return nil
}
