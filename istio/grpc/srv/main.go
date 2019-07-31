package main

import (
	"flag"

	grpcClient "github.com/hb-go/micro-plugins/client/istio_grpc"
	grpcServer "github.com/hb-go/micro-plugins/server/istio_grpc"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"

	"github.com/hb-go/micro/istio/grpc/srv/handler"
	example "github.com/hb-go/micro/istio/grpc/srv/proto/example"
	"github.com/hb-go/micro/istio/grpc/srv/subscriber"
)

var (
	serverAddr string
	callAddr   string
	cmdHelp    bool
)

func init() {
	flag.StringVar(&serverAddr, "server_address", "0.0.0.0:9080", "server address.")
	flag.StringVar(&callAddr, "client_call_address", ":9080", "client call options address.")
	flag.BoolVar(&cmdHelp, "h", false, "help")
	flag.Parse()
}

func main() {
	if cmdHelp {
		flag.PrintDefaults()
		return
	}

	c := grpcClient.NewClient(
		client.ContentType("application/json"),
		func(o *client.Options) {
			o.CallOptions.Address = callAddr
		},
	)
	s := grpcServer.NewServer(
		server.Address(serverAddr),
	)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.sample"),
		micro.Version("latest"),
		micro.Client(c),
		micro.Server(s),

		// 兼容micro cmd parse
		micro.Flags(cli.StringFlag{
			Name:   "client_call_address",
			EnvVar: "MICRO_CLIENT_CALL_ADDRESS",
			Usage:  " Invalid!!!",
		}),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("topic.go.micro.srv.sample", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("topic.go.micro.srv.sample", service.Server(), subscriber.Handler)

	// Initialise service
	service.Init()

	log.Logf("Service Run")

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
