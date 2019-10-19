package handler

import (
	"context"
	"encoding/json"

	"github.com/micro/go-micro/util/log"

	"github.com/hb-go/micro/console/api/client"
	user "github.com/hb-go/micro/console/srv/proto/user"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
)

type User struct{}

type responseBody struct {
	Code   int64       `json:"code"`
	Detail string      `json:"detail"`
	Data   interface{} `json:"data"`
}

func ResponseBody(code int64, data interface{}, detail ...string) (string, error) {
	body := responseBody{
		Code: code,
		Data: data,
	}

	if len(detail) > 0 {
		body.Detail = detail[0]
	}

	b, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// Example.Call is called by the API as /example/call with post body {"name": "foo"}
func (e *User) Login(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Example.Call request")

	// extract the client from the context
	userClient, ok := client.UserFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.example.example.call", "example client not found")
	}

	// make request
	r := &user.LoginRequest{}
	if err := json.Unmarshal([]byte(req.GetBody()), r); err != nil {
		return err
	}

	response, err := userClient.Login(ctx, r)
	if err != nil {
		return errors.InternalServerError("go.micro.api.example.example.call", err.Error())
	}

	b, err := ResponseBody(20000, response)
	log.Logf("err:%v", err)
	if err != nil {
		return errors.InternalServerError("go.micro.api.example.example.call", err.Error())
	}
	log.Log(b)
	rsp.StatusCode = 200
	rsp.Body = b

	return nil
}

// Example.Call is called by the API as /example/call with post body {"name": "foo"}
func (e *User) Logout(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Example.Call request")

	// extract the client from the context
	userClient, ok := client.UserFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.example.example.call", "example client not found")
	}

	// make request
	response, err := userClient.Logout(ctx, &user.Request{
		Id: "admin",
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.example.example.call", err.Error())
	}

	b, err := ResponseBody(20000, response)
	if err != nil {
		return errors.InternalServerError("go.micro.api.example.example.call", err.Error())
	}

	rsp.StatusCode = 200
	rsp.Body = b

	return nil
}

// Example.Call is called by the API as /example/call with post body {"name": "foo"}
func (e *User) Info(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Example.Call request")

	// extract the client from the context
	userClient, ok := client.UserFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.example.example.call", "example client not found")
	}

	// make request
	response, err := userClient.Info(ctx, &user.Request{
		Id: "admin",
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.example.example.call", err.Error())
	}

	b, err := ResponseBody(20000, response)
	if err != nil {
		return errors.InternalServerError("go.micro.api.example.example.call", err.Error())
	}

	rsp.StatusCode = 200
	rsp.Body = b

	return nil
}
