package main

import (
	"github.com/hb-go/micro/console/srv/handler"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	user "github.com/hb-go/micro/console/srv/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.console"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
