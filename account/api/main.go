package main

import (
	"github.com/micro/go-log"

	"github.com/micro/go-micro"
	"github.com/hb-go/micro/account/api/handler"
	"github.com/hb-go/micro/account/api/client"

	example "github.com/hb-go/micro/account/api/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.account"),
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
