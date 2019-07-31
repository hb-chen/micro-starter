package main

import (
	"context"
	"flag"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-plugins/codec/bsonrpc"
	"github.com/micro/go-plugins/codec/jsonrpc2"
	"github.com/micro/go-plugins/codec/msgpackrpc"
	"github.com/micro/go-plugins/transport/tcp"

	"github.com/hb-go/micro/benchmark/service"
)

var concurrency = flag.Int("c", 1, "concurrency")
var total = flag.Int("n", 1, "total requests for all clients")
var contentType = flag.String("ct", "protobuf", "content type")

// "application/grpc":         grpc.NewCodec,           // TPS:9990
// "application/protobuf":     proto.NewCodec,          // TPS:11037
// "application/json":         json.NewCodec,           // TPS:7782
// "application/json-rpc":     jsonrpc.NewCodec,        // TPS:6544
// "application/proto-rpc":    protorpc.NewCodec,       // TPS:10869
// "application/octet-stream": raw.NewCodec,            // error:{"id":"go.micro.client.codec","code":500,"detail":"failed to write: field1:
// "application/msgpackrpc", msgpackrpc.NewCodec,       // 需要实现EncodeMsg(*Writer) error，error:{"id":"go.micro.client.codec","code":500,"detail":"Not encodable","status":"Internal Server Error"}
// "application/bsonrpc", bsonrpc.NewCodec,             // TPS:4970
// "application/jsonrpc2", jsonrpc2.NewCodec,           // error:{"id":"go.micro.client.transport","code":500,"detail":"EOF","status":"Internal Server Error"}

func main() {
	flag.Parse()
	serviceName := "hello.tcp.rpc"
	service.ClientWith(
		serviceName,
		func() client.Client {
			cache := selector.NewSelector(func(o *selector.Options) {
				o.Context = context.WithValue(o.Context, "selector_ttl", time.Minute*3)
			})
			return client.NewClient(
				client.Codec("application/msgpackrpc", msgpackrpc.NewCodec),
				client.Codec("application/bsonrpc", bsonrpc.NewCodec),
				client.Codec("application/jsonrpc2", jsonrpc2.NewCodec),
				client.Transport(tcp.NewTransport()),
				client.ContentType("application/"+*contentType),
				client.Selector(cache),
				client.Retries(1),
				client.PoolSize(*concurrency*2),
				client.RequestTimeout(time.Millisecond*100),
				// client.Wrap(breaker.NewClientWrapper()),
				// client.Wrap(ratelimit.NewClientWrapper(10000)),
			)
		},
		*concurrency,
		*total,
	)

}
