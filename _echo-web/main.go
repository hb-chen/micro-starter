package main

import (
	"time"

	"github.com/micro/go-log"
	"github.com/micro/go-web"

	"github.com/hb-go/micro/_echo-web/router"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.echo"),
		web.Version("latest"),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
	)

	service.Handle("/", router.NewRouter())

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
