package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/micro/go-log"

	"github.com/micro/go-micro/client"
	example "github.com/hb-go/micro/post/srv/proto/example"


	"golang.org/x/net/context"
)

func ExampleCall(w http.ResponseWriter, r *http.Request) {
	log.Log("Req ExampleCall")

	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Log("decode error!")
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	exampleClient := example.NewExampleService("go.micro.srv.post", client.DefaultClient)
	rsp, err := exampleClient.Call(context.TODO(), &example.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		log.Log("example call error!")
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Log("encode error!")
		http.Error(w, err.Error(), 500)
		return
	}
}
