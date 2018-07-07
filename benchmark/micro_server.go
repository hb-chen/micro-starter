package main

import (
	"flag"
	fmt "fmt"
	"time"

	"golang.org/x/net/context"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-plugins/transport/tcp"
	"github.com/micro/go-plugins/wrapper/ratelimiter/uber"

	"github.com/hb-go/micro/benchmark/proto"
)

type HelloS struct{}

func (t *HelloS) Say(ctx context.Context, args *proto.BenchmarkMessage, reply *proto.BenchmarkMessage) error {
	s := "OK"
	var i int32 = 100
	args.Field1 = &s
	args.Field2 = &i
	*reply = *args
	if *delay > 0 {
		time.Sleep(*delay)
	}
	return nil
}

//var host = flag.String("s", "127.0.0.1:8972", "listened ip and port")

var delay = flag.Duration("delay", 0, "delay to mock business processing")

func main() {
	flag.Parse()

	service := micro.NewService(
		micro.Name("hello"),
		micro.Version("latest"),
		micro.Transport(tcp.NewTransport()),
	)

	service.Init(
		// handler wrap
		micro.WrapHandler(
			ratelimit.NewHandlerWrapper(100),
		),
	)

	proto.RegisterHelloHandler(service.Server(), &HelloS{})

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
