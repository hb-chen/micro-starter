package main

import (
	"time"

	grpcClient "github.com/hb-go/micro-plugins/client/istio_grpc"
	"github.com/micro/go-api/proto"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"

	example "github.com/hb-go/micro/istio/grpc/api/proto/example"
)

const (
	address = "localhost:8082"
)

func main() {
	c := grpcClient.NewClient(
		client.ContentType("application/json"),
		func(o *client.Options) {
			o.CallOptions.Address = address
		},
	)
	client := example.NewExampleService("go.micro.api.sample", c)

	post := make(map[string]*go_api.Pair)
	post["name"] = &go_api.Pair{
		Key:    "name",
		Values: []string{"Hobo"},
	}

	req := &go_api.Request{
		Method: "POST",
		Post:   post,
	}

	now := time.Now()
	resp, err := client.Call(context.Background(), req)
	if err != nil {
		log.Logf("did not connect: %v", err)
	}

	log.Logf("resp: %v", resp)
	log.Logf("duration: %s", time.Since(now).String())
}
