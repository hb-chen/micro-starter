package main

import (
	"github.com/micro/go-log"

	"github.com/micro/go-micro"
	"github.com/hb-go/micro/post/api/handler"
	"github.com/hb-go/micro/post/api/client"

	example "github.com/hb-go/micro/post/api/proto/example"
	post "github.com/hb-go/micro/post/api/proto/post"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.post"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	post.RegisterPostHandler(service.Server(), new(handler.Post))

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
