package main

import (
	_ "github.com/micro/go-plugins/registry/kubernetes"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/hb-go/micro/auth/srv/handler"
	"github.com/hb-go/micro/auth/srv/subscriber"

	example "github.com/hb-go/micro/auth/srv/proto/example"
	token "github.com/hb-go/micro/auth/srv/proto/token"
	user "github.com/hb-go/micro/auth/srv/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.auth"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("topic.go.micro.srv.auth", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("topic.go.micro.srv.auth", service.Server(), subscriber.Handler)

	// Token Handler
	token.RegisterTokenHandler(service.Server(), new(handler.Token))

	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
