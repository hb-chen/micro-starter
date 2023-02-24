package server

import (
	"github.com/micro/micro/v3/service/server"
	"go.uber.org/dig"

	pb "github.com/hb-chen/micro-starter/service/greeting/proto/greeting"
)

func Apply(server server.Server, c *dig.Container) {
	c.Invoke(func(svc pb.GreetingHandler) {
		pb.RegisterGreetingHandler(server, svc)
	})
}
