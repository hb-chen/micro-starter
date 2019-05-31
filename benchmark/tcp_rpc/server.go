package main

import (
	"flag"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/codec/bsonrpc"
	"github.com/micro/go-plugins/codec/jsonrpc2"
	"github.com/micro/go-plugins/codec/msgpackrpc"
	"github.com/micro/go-plugins/transport/tcp"

	"github.com/hb-go/micro/benchmark/service"
)

var delay = flag.Duration("delay", 0, "delay to mock business processing")

func main() {
	flag.Parse()

	serviceName := "hello.tcp.rpc"
	service.ServeWith(
		serviceName,
		micro.Server(
			server.NewServer(
				server.Codec("application/msgpackrpc", msgpackrpc.NewCodec),
				server.Codec("application/bsonrpc", bsonrpc.NewCodec),
				server.Codec("application/jsonrpc2", jsonrpc2.NewCodec),
			)),
		micro.Transport(tcp.NewTransport()),
		*delay,
	)
}
