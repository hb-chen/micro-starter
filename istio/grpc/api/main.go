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

	apiClient "github.com/hb-go/micro/istio/grpc/api/client"
	"github.com/hb-go/micro/istio/grpc/api/handler"
	example "github.com/hb-go/micro/istio/grpc/api/proto/example"
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
		micro.Name("go.micro.api.sample"),
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

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		micro.WrapHandler(apiClient.ExampleWrapper(service)),
	)

	log.Logf("Service Run")

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
