package handler

import (
	"github.com/micro/micro/v3/service/server"
	"go.uber.org/dig"

	pb "github.com/hb-chen/micro-starter/service/account/proto/account"
	"github.com/hb-chen/micro-starter/service/account/usecase"
)

func Apply(server server.Server, c *dig.Container) {
	c.Invoke(func(userUseCase usecase.UserUseCase) {
		pb.RegisterAccountHandler(server, NewAccountService(userUseCase))
	})
}
