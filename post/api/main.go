package main

import (
	"github.com/micro/go-log"

	"github.com/micro/go-micro"
	"github.com/hb-go/micro/post/api/handler"
	"github.com/hb-go/micro/post/api/client"

	example "github.com/hb-go/micro/post/api/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.post"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		micro.WrapHandler(client.ExampleWrapper(service)),
	)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
