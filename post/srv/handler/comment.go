package handler

import (
	"github.com/micro/go-log"

	comment "github.com/hb-go/micro/post/srv/proto/comment"
	"golang.org/x/net/context"
)

type Comment struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Comment) Call(ctx context.Context, req *comment.Request, rsp *comment.Response) error {
	log.Log("Received Comment.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Comment) Stream(ctx context.Context, req *comment.StreamingRequest, stream comment.Comment_StreamStream) error {
	log.Logf("Received Comment.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&comment.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Comment) PingPong(ctx context.Context, stream comment.Comment_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&comment.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
