package server

import (
	log "github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/server"
	"go.uber.org/dig"

	pb "github.com/hb-chen/micro-starter/service/greeting/proto/greeting"
)

func Apply(server server.Server, c *dig.Container) {
	err := c.Invoke(func(svc pb.GreetingHandler) {
		err := pb.RegisterGreetingHandler(server, svc)
		if err != nil {
			log.Error(err)
			return
		}
	})
	if err != nil {
		log.Error(err)
		return
	}
}
