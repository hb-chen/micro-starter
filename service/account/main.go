package main

import (
	"github.com/hb-chen/micro-starter/service/account/interface/handler"
	"github.com/hb-chen/micro-starter/service/account/registry"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// New Service
	srv := service.New(
		service.Name("account"),
		service.Version("latest"),
	)

	c, err := registry.NewContainer()
	if err != nil {
		logger.Fatal(err)
	}

	// Register Handler
	handler.Apply(srv.Server(), c)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
