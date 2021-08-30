package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	"github.com/hb-chen/micro/console/srv/handler"
	pb "github.com/hb-chen/micro/console/srv/proto/user"
)

func main() {
	// New Service
	srv := service.New(
		service.Name("console"),
		service.Version("latest"),
	)

	// Register Handler
	pb.RegisterUserHandler(srv.Server(), new(handler.User))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
