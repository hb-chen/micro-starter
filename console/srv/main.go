package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	"github.com/hb-chen/micro-starter/console/srv/handler"
	pb "github.com/hb-chen/micro-starter/console/srv/proto/account"
)

func main() {
	// New Service
	srv := service.New(
		service.Name("console"),
		service.Version("latest"),
	)

	// Register Handler
	pb.RegisterAccountHandler(srv.Server(), new(handler.Account))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
