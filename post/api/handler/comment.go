package handler

import (
	"encoding/json"

	"github.com/micro/go-log"

	"github.com/hb-go/micro/post/api/client"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/go-api/proto"
	comment "github.com/hb-go/micro/post/api/proto/comment"
	postSrv "github.com/hb-go/micro/post/srv/proto/post"

	"golang.org/x/net/context"
)

type Comment struct{}

// Post.Comments is called by the API as /post/comment/comments with post body {"name": "foo"}
func (e *Comment) Comments(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Post.Comments request")

	// extract the client from the context

	id, ok := req.Get["id"]
	log.Log("Req post comments id:%v", id.String())

	response := comment.Rsp{}

	commentClient, ok := client.CommentFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.post.comment", "post client not found")
	}

	// make request
	rspComments, err := commentClient.GetComments(ctx, &postSrv.Req{
		Id: 0,
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.post.comment.GetComments", err.Error())
	}
	response.Comments = rspComments.Comments

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
