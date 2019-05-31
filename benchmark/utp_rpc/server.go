package main

import (
	"flag"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/transport/utp"

	"github.com/hb-go/micro/benchmark/service"
)

var delay = flag.Duration("delay", 0, "delay to mock business processing")

func main() {
	flag.Parse()

	serviceName := "hello.utp.rpc"
	service.ServeWith(
		serviceName,
		micro.Server(server.DefaultServer),
		micro.Transport(utp.NewTransport()),
		*delay,
	)
}
