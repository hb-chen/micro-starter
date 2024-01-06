package main

import (
	"micro.dev/v4/service"
	"micro.dev/v4/service/logger"

	svc "github.com/hb-chen/micro-starter/pkg/service"
	_ "github.com/hb-chen/micro-starter/profile"
	"github.com/hb-chen/micro-starter/service/greeting/registry"
	"github.com/hb-chen/micro-starter/service/greeting/server"
)

func main() {
	// New Service
	srv := svc.New(
		service.Name("greeting"),
		service.Version("latest"),
	)

	c, err := registry.NewContainer()
	if err != nil {
		logger.Fatal(err)
	}

	// Register Handler
	server.Apply(srv.Server(), c)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
