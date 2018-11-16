package main

import (
	"flag"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	
	httpClient "github.com/hb-go/micro-plugins/client/istio_http"
	httpServer "github.com/hb-go/micro-plugins/server/istio_http"
	"github.com/hb-go/micro/istio/http/srv/handler"
	example "github.com/hb-go/micro/istio/http/srv/proto/example"
	"github.com/micro/go-plugins/registry/noop"
)

var (
	serverAddr string
	callAddr   string
	cmdHelp    bool
)

// TODO 命令参数与micro的兼容
func init() {
	flag.StringVar(&serverAddr, "sa", "0.0.0.0:9080", "server address.")
	flag.StringVar(&callAddr, "ca", ":9080", "client call options address.")
	flag.BoolVar(&cmdHelp, "h", false, "help")
	flag.Parse()
}

func main() {
	if cmdHelp {
		flag.PrintDefaults()
		return
	}
	
	// 多client需要统一端口，或者在client中hard code
	c := httpClient.NewClient(
		client.ContentType("application/json"),
		func(o *client.Options) {
			o.CallOptions.Address = callAddr
		},
	)
	s := httpServer.NewServer(
		server.Address(serverAddr),
	)
	
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.sample"),
		micro.Version("latest"),
		micro.Registry(noop.NewRegistry()),
		micro.Client(c),
		micro.Server(s),
	)
	
	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))
	
	// Register Struct as Subscriber
	// micro.RegisterSubscriber("topic.go.micro.srv.sample", service.Server(), new(subscriber.Example))
	
	// Register Function as Subscriber
	// micro.RegisterSubscriber("topic.go.micro.srv.sample", service.Server(), subscriber.Handler)
	
	// Initialise service
	service.Init()
	
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
