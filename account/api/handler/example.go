package handler

import (
	"encoding/json"
	"github.com/micro/go-log"

	"github.com/hb-go/micro/account/api/client"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/go-api/proto"
	example "github.com/micro/examples/template/srv/proto/example"

	"golang.org/x/net/context"
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

// Example.Call is called by the API as /account/example/call with post body {"name": "foo"}
func (e *Example) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Example.Call request")

	// extract the client from the context
	exampleClient, ok := client.ExampleFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.account.example.call", "example client not found")
	}

	// make request
	response, err := exampleClient.Call(ctx, &example.Request{
		Name: extractValue(req.Post["name"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.account.example.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
