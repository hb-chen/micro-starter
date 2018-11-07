package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"

	httpClient "github.com/hb-go/micro-plugins/client/istio_http"
	httpServer "github.com/hb-go/micro-plugins/server/istio_http"
	apiClient "github.com/hb-go/micro/istio/http/api/client"
	"github.com/hb-go/micro/istio/http/api/handler"
	example "github.com/hb-go/micro/istio/http/api/proto/example"
	"github.com/micro/go-plugins/registry/noop"
)

func main() {
	c := httpClient.NewClient(
		client.ContentType("application/json"),
	)
	s := httpServer.NewServer(
		server.Address("localhost:8081"),
	)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.http"),
		micro.Version("latest"),
		micro.Registry(noop.NewRegistry()),
		micro.Client(c),
		micro.Server(s),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		micro.WrapHandler(apiClient.ExampleWrapper(service)),
	)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
