package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/hb-go/micro/post/srv/handler"
	"github.com/hb-go/micro/post/srv/subscriber"

	example "github.com/hb-go/micro/post/srv/proto/example"
	post "github.com/hb-go/micro/post/srv/proto/post"
	comment "github.com/hb-go/micro/post/srv/proto/comment"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.post"),
		micro.Version("latest"),
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
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
