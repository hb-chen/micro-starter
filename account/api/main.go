package main

import (
	"context"
	"time"

	"github.com/micro/go-api"
	ha "github.com/micro/go-api/handler/api"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	// "github.com/micro/go-plugins/wrapper/trace/opentracing"
	breaker "github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"github.com/micro/go-plugins/wrapper/ratelimiter/uber"

	"github.com/hb-go/micro/account/api/client"
	"github.com/hb-go/micro/pkg/wrapper/auth"
	// tracer "github.com/hb-go/micro/pkg/opentracing"
	"github.com/hb-go/micro/account/api/handler"
	account "github.com/hb-go/micro/account/api/proto/account"
	example "github.com/hb-go/micro/account/api/proto/example"
)

func main() {
	// Tracer
	// t, closer, err := tracer.NewJaegerTracer("account.api", "127.0.0.1:6831")
	// if err != nil {
	// 	log.Fatalf("opentracing tracer create error:%v", err)
	// }
	// defer closer.Close()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.account"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		// micro.WrapClient(opentracing.NewClientWrapper(t)),
		// micro.WrapHandler(opentracing.NewHandlerWrapper(t)),
	)

	// graceful
	service.Server().Init(
		server.Wait(nil),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example),
		api.WithEndpoint(&api.Endpoint{
			// The RPC method
			Name: "Example.Call",
			// The HTTP paths. This can be a POSIX regex
			Path: []string{"/example/call"},
			// The HTTP Methods for this endpoint
			Method: []string{"GET"},
			// The API handler to use
			Handler: ha.Handler,
		}))

	account.RegisterAccountHandler(service.Server(), new(handler.Account),
		api.WithEndpoint(&api.Endpoint{
			// The RPC method
			Name: "Account.Login",
			// The HTTP paths. This can be a POSIX regex
			Path: []string{"/login"},
			// The HTTP Methods for this endpoint
			Method: []string{"POST"},
			// The API handler to use
			Handler: ha.Handler,
		}),
		api.WithEndpoint(&api.Endpoint{
			// The RPC method
			Name: "Account.Register",
			// The HTTP paths. This can be a POSIX regex
			Path: []string{"/register"},
			// The HTTP Methods for this endpoint
			Method: []string{"POST"},
			// The API handler to use
			Handler: ha.Handler,
		}))

	// Initialise service
	service.Init(
		// client wrap
		micro.WrapClient(
			// @TODO fallback、hystrix.ConfigureCommand()
			breaker.NewClientWrapper(),
			ratelimit.NewClientWrapper(1024),
		),
		// handler wrap
		micro.WrapHandler(
			ratelimit.NewHandlerWrapper(1024),
			auth.NewHandlerWrapper(
				service,
				auth.ServiceName("go.micro.srv.auth"),
				auth.Skipper(authSkipperFunc),
			),
			client.ExampleWrapper(service),
			client.UserWrapper(service),
			client.TokenWrapper(service),
		),
	)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func authSkipperFunc(ctx context.Context, req server.Request) bool {
	skipMethods := map[string]bool{
		"Account.Login":    true,
		"Account.Register": false,
	}

	log.Logf("req Service:%v", req.Service())
	log.Logf("req Method:%v", req.Method())
	log.Logf("req Endpoint:%v", req.Endpoint())
	log.Logf("req ContentType:%v", req.ContentType())
	log.Logf("req Header:%v", req.Header())

	if skip, ok := skipMethods[req.Method()]; ok && skip {
		return true
	}

	return false
}
