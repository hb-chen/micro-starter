package server

import (
	"go.uber.org/dig"
	log "micro.dev/v4/service/logger"
	"micro.dev/v4/service/server"

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
