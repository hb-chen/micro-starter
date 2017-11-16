package main

import (
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-os/trace"
	"github.com/micro/go-plugins/trace/zipkin"

	"github.com/hb-go/micro/account/api/handler"
	"github.com/hb-go/micro/account/api/client"
	account "github.com/hb-go/micro/account/api/proto/account"
	example "github.com/hb-go/micro/account/api/proto/example"
)

func main() {
	t := zipkin.NewTrace(
		trace.Topic("zipkin"),
		trace.Collectors("localhost:9092"),
	)
	defer t.Close()

	srv := &registry.Service{Name: "go.micro.api.account"}

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.account"),
		micro.Version("latest"),
		micro.WrapClient(trace.ClientWrapper(t, srv)),
		micro.WrapHandler(trace.HandlerWrapper(t, srv)),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	account.RegisterAccountHandler(service.Server(), new(handler.Account))

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		micro.WrapHandler(
			client.ExampleWrapper(service),
			client.UserWrapper(service),
			client.TokenWrapper(service),
		),
	)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
