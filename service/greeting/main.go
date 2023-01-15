package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	_ "github.com/hb-chen/micro-starter/profile"
	"github.com/hb-chen/micro-starter/service/greeting/interface/handler"
	"github.com/hb-chen/micro-starter/service/greeting/registry"
)

func main() {
	// New Service
	srv := service.New(
		service.Name("greeting"),
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
