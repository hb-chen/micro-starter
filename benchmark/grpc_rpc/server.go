package main

import (
	"flag"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/transport/grpc"

	"github.com/hb-go/micro/benchmark/service"
)

var delay = flag.Duration("delay", 0, "delay to mock business processing")

func main() {
	flag.Parse()

	serviceName := "hello.grpc.rpc"
	service.ServeWith(
		serviceName,
		micro.Server(server.DefaultServer),
		micro.Transport(grpc.NewTransport()),
		*delay,
	)
}
