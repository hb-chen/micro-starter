package handler

import (
	"github.com/micro/go-log"

	post "github.com/hb-go/micro/post/srv/proto/post"
	comment "github.com/hb-go/micro/post/srv/proto/comment"
	"golang.org/x/net/context"
)

type Comment struct{}

// Call is a single request handler called via client.GetComments or the generated client code
func (e *Comment) GetComments(ctx context.Context, req *post.Req, rsp *comment.Rsp) error {
	log.Log("Received Comment.GetComments request")

	rsp.Comments = append(rsp.Comments, &comment.CommentDto{Content:"content"})
	rsp.Comments = append(rsp.Comments, &comment.CommentDto{Content:"content"})

	return nil
}