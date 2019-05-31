package service

import (
	"time"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"

	"github.com/hb-go/micro/benchmark/proto"
)

type HelloS struct{}

func (t *HelloS) Say(ctx context.Context, args *proto.BenchmarkMessage, reply *proto.BenchmarkMessage) error {
	s := "OK"
	var i int32 = 100
	args.Field1 = &s
	args.Field2 = &i
	*reply = *args
	if delay > 0 {
		time.Sleep(delay)
	}
	return nil
}

var delay = time.Duration(0)

func ServeWith(serviceName string, serverOpt, transOpt micro.Option, d time.Duration) {
	delay = d
	service := micro.NewService(
		serverOpt,
		micro.Name(serviceName),
		micro.Version("latest"),
		transOpt,
	)

	service.Init(
		// handler wrap
		// micro.WrapHandler(
		// 	ratelimit.NewHandlerWrapper(10000),
		// ),
	)

	proto.RegisterHelloHandler(service.Server(), &HelloS{})

	log.Logf("service: %s start", serviceName)

	// Run the server
	if err := service.Run(); err != nil {
		log.Log(err)
	}
}
