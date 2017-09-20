package handler

import (
	"github.com/micro/go-log"

	post "github.com/hb-go/micro/post/srv/proto/post"
	"golang.org/x/net/context"
)

type Post struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Post) Call(ctx context.Context, req *post.Request, rsp *post.Response) error {
	log.Log("Received Post.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Post) Stream(ctx context.Context, req *post.StreamingRequest, stream post.Post_StreamStream) error {
	log.Logf("Received Post.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&post.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Post) PingPong(ctx context.Context, stream post.Post_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&post.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
