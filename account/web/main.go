package main

import (
        "github.com/micro/go-log"
	"net/http"

        "github.com/micro/go-web"
        "github.com/hb-go/micro/account/web/handler"
)

func main() {
	// create new web service
        service := web.NewService(
                web.Name("go.micro.web.account"),
                web.Version("latest"),
        )

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/example/call", handler.ExampleCall)

	// initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
