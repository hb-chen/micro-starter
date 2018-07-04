package main

import (
	"time"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/micro/go-plugins/wrapper/ratelimiter/uber"

	tracer "github.com/hb-go/micro/pkg/opentracing"
	"github.com/hb-go/micro/auth/srv/handler"
	"github.com/hb-go/micro/auth/srv/subscriber"
	example "github.com/hb-go/micro/auth/srv/proto/example"
	token "github.com/hb-go/micro/auth/srv/proto/token"
	user "github.com/hb-go/micro/auth/srv/proto/user"
)

func main() {
	// Tracer
	t, closer, err := tracer.NewJaegerTracer("auth.srv", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.auth"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.WrapClient(opentracing.NewClientWrapper(t)),
		micro.WrapHandler(opentracing.NewHandlerWrapper(t)),
	)

	// graceful
	service.Server().Init(
		server.Wait(true),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("topic.go.micro.srv.auth", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("topic.go.micro.srv.auth", service.Server(), subscriber.Handler)

	// Token Handler
	token.RegisterTokenHandler(service.Server(), new(handler.Token))

	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Initialise service
	service.Init(
		// handler wrap
		micro.WrapHandler(
			ratelimit.NewHandlerWrapper(1024),
		),
	)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
