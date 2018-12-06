package handler

import (
	"encoding/json"
	"sync"

	"github.com/micro/go-log"
	"golang.org/x/net/context"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/go-api/proto"

	"github.com/hb-go/micro/istio/grpc/api/client"
	example "github.com/hb-go/micro/istio/grpc/srv/proto/example"
)

type Example struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// Example.Call is called by the API as /grpc/example/call with post body {"name": "foo"}
func (e *Example) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Example.Call request")

	// extract the client from the context
	exampleClient, ok := client.ExampleFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.sample.example.call", "example client not found")
	}

	// make request
	response, err := exampleClient.Call(ctx, &example.Request{
		Name: extractValue(req.Post["name"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.sample.example.call", err.Error())
	}

	b, _ := json.Marshal(response)

	// stream
	count := int64(3)
	stream, err := exampleClient.Stream(ctx, &example.StreamingRequest{Count: count})
	if err != nil {
		return errors.InternalServerError("go.micro.api.sample.example.stream", err.Error())
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for i := int64(0); i < count; i++ {
			streamRsp, err := stream.Recv()
			if err != nil {
				log.Logf("stream recv error: %v", err.Error())
			} else {
				log.Logf("stream recv count: %d", streamRsp.Count)
			}
		}
		stream.Close()

		wg.Done()
	}()

	wg.Wait()

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
