package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/micro/go-log"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/selector/cache"

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

	time.Sleep(time.Millisecond *5)

	client.NewClient(
		client.Selector(cache.NewSelector(cache.TTL(time.Second *30))),
		client.Transport(tcp.NewTransport()),
	)

	// call the backend service
	userClient := user.NewUserClient("go.micro.srv.auth", client.DefaultClient)
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
