package main

import (
	"time"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-api"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"

	tracer "github.com/hb-go/micro/pkg/opentracing"
	"github.com/hb-go/micro/post/api/handler"
	"github.com/hb-go/micro/post/api/client"
	example "github.com/hb-go/micro/post/api/proto/example"
	post "github.com/hb-go/micro/post/api/proto/post"
	comment "github.com/hb-go/micro/post/api/proto/comment"
)

func main() {
	// Tracer
	t, closer, err := tracer.NewJaegerTracer("post.srv", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.post"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.WrapClient(opentracing.NewClientWrapper(t)),
		micro.WrapHandler(opentracing.NewHandlerWrapper(t)),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example),
		api.WithEndpoint(&api.Endpoint{
			// The RPC method
			Name: "Example.Call",
			// The HTTP paths. This can be a POSIX regex
			Path: []string{"/example/call"},
			// The HTTP Methods for this endpoint
			Method: []string{"GET"},
			// The API handler to use
			Handler: api.Api,
		}))

	post.RegisterPostHandler(service.Server(), new(handler.Post),
		api.WithEndpoint(&api.Endpoint{
			// The RPC method
			Name: "Post.Post",
			// The HTTP paths. This can be a POSIX regex
			Path: []string{"/post"},
			// The HTTP Methods for this endpoint
			Method: []string{"GET"},
			// The API handler to use
			Handler: api.Api,
		}))

	comment.RegisterCommentHandler(service.Server(), new(handler.Comment),
		api.WithEndpoint(&api.Endpoint{
			// The RPC method
			Name: "Comment.Comments",
			// The HTTP paths. This can be a POSIX regex
			Path: []string{"/post/comments"},
			// The HTTP Methods for this endpoint
			Method: []string{"GET"},
			// The API handler to use
			Handler: api.Api,
		}))

	// Initialise service
	service.Init(
		// create wraps for the srv clients
		micro.WrapHandler(
			client.ExampleWrapper(service),
			client.PostWrapper(service),
			client.CommentWrapper(service),
		),
	)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
