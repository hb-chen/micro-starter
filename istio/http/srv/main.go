package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"

	httpClient "github.com/hb-go/micro-plugins/client/istio_http"
	httpServer "github.com/hb-go/micro-plugins/server/istio_http"
	"github.com/hb-go/micro/istio/http/srv/handler"
	example "github.com/hb-go/micro/istio/http/srv/proto/example"
	"github.com/micro/go-plugins/registry/noop"
)

func main() {
	c := httpClient.NewClient(
		client.ContentType("application/json"),
	)
	s := httpServer.NewServer(
		server.Address("localhost:8082"),
	)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.test"),
		micro.Version("latest"),
		micro.Registry(noop.NewRegistry()),
		micro.Client(c),
		micro.Server(s),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("topic.go.micro.srv.http", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("topic.go.micro.srv.http", service.Server(), subscriber.Handler)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
