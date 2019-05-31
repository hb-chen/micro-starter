package main

import (
	"flag"

	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/server/grpc"
	"github.com/micro/go-plugins/transport/tcp"

	"github.com/hb-go/micro/benchmark/service"
)

var delay = flag.Duration("delay", 0, "delay to mock business processing")

func main() {
	flag.Parse()

	serviceName := "hello.utp.grpc"
	service.ServeWith(
		serviceName,
		micro.Server(grpc.NewServer(
			grpc.Codec("application/test", service.ProtoCodec{}), // gRPC自定义codec测试
		)),
		micro.Transport(tcp.NewTransport()),
		*delay,
	)
}
