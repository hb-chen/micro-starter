package main

import (
	"time"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/micro/go-plugins/wrapper/ratelimiter/uber"

	tracer "github.com/hb-go/micro/pkg/opentracing"
	"github.com/hb-go/micro/post/srv/handler"
	"github.com/hb-go/micro/post/srv/subscriber"
	example "github.com/hb-go/micro/post/srv/proto/example"
	post "github.com/hb-go/micro/post/srv/proto/post"
	comment "github.com/hb-go/micro/post/srv/proto/comment"
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
		micro.Name("go.micro.srv.post"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.WrapClient(opentracing.NewClientWrapper(t)),
		micro.WrapHandler(opentracing.NewHandlerWrapper(t)),
	)

	// graceful
	service.Server().Init(
		server.Wait(true),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("topic.go.micro.srv.post", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("topic.go.micro.srv.post", service.Server(), subscriber.Handler)

	// Post
	post.RegisterPostHandler(service.Server(), new(handler.Post))

	// Comment
	comment.RegisterCommentHandler(service.Server(), new(handler.Comment))

	// Initialise service
	service.Init(
		// handler wrap
		micro.WrapHandler(
			ratelimit.NewHandlerWrapper(1024),
		),
	)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
