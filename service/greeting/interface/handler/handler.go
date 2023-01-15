package handler

import (
	"github.com/micro/micro/v3/service/server"
	"go.uber.org/dig"

	pb "github.com/hb-chen/micro-starter/service/greeting/proto/greeting"
	"github.com/hb-chen/micro-starter/service/greeting/usecase"
)

func Apply(server server.Server, c *dig.Container) {
	c.Invoke(func(useCase usecase.GreetingUseCase) {
		pb.RegisterGreetingHandler(server, NewGreetingService(useCase))
	})
}
