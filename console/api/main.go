package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/api"
	ha "github.com/micro/go-micro/api/handler/api"
	"github.com/micro/go-micro/util/log"

	"github.com/hb-go/micro/console/api/client"
	"github.com/hb-go/micro/console/api/handler"

	user "github.com/hb-go/micro/console/api/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.console"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		micro.WrapHandler(client.UserWrapper(service)),
	)

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User),
		api.WithEndpoint(&api.Endpoint{
			// The RPC method
			Name: "User.Login",
			// The HTTP paths. This can be a POSIX regex
			Path: []string{"/user/login"},
			// The HTTP Methods for this endpoint
			Method: []string{"POST", "OPTIONS"},
			// The API handler to use
			Handler: ha.Handler,
		}),
		api.WithEndpoint(&api.Endpoint{
			// The RPC method
			Name: "User.Logout",
			// The HTTP paths. This can be a POSIX regex
			Path: []string{"/user/logout"},
			// The HTTP Methods for this endpoint
			Method: []string{"POST"},
			// The API handler to use
			Handler: ha.Handler,
		}),
		api.WithEndpoint(&api.Endpoint{
			// The RPC method
			Name: "User.Info",
			// The HTTP paths. This can be a POSIX regex
			Path: []string{"/user/info"},
			// The HTTP Methods for this endpoint
			Method: []string{"GET"},
			// The API handler to use
			Handler: ha.Handler,
		}),
	)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
