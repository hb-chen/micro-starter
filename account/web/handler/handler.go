package handler

import (
	"encoding/json"
	"net/http"
	"time"
	"context"

	"github.com/micro/go-log"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/selector"

	user "github.com/hb-go/micro/auth/srv/proto/user"
	"github.com/micro/go-plugins/transport/tcp"
)

func ExampleCall(w http.ResponseWriter, r *http.Request) {
	log.Logf("example call")
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Logf("example call decode err!")
		http.Error(w, err.Error(), 500)
		return
	}

	time.Sleep(time.Millisecond * 5)

	cache := selector.NewSelector(func(o *selector.Options) {
		o.Context = context.WithValue(o.Context, "selector_ttl", time.Second*30)
	})
	client.NewClient(
		client.Selector(cache),
		client.Transport(tcp.NewTransport()),
	)

	// call the backend service
	userClient := user.NewUserService("go.micro.srv.auth", client.DefaultClient)
	rsp, err := userClient.GetUserLogin(context.TODO(), &user.ReqLogin{
		Nickname: "Hobo",
		Pwd:      "pwd",
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"user": rsp,
		"ref":  time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Logf("example call encode err!")
		http.Error(w, err.Error(), 500)
		return
	}
}
