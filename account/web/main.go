package main

import (
	"net/http"
	"time"

	"github.com/micro/go-log"
	"github.com/micro/go-web"

	"github.com/hb-go/micro/account/web/handler"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.account"),
		web.Version("latest"),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
	)

	// register html handler
	// 注意Dir的相对路径
	// 在web目录下go run main.go http.Dir("html")
	// 在micro目录下go run account/web/main.go http.Dir("account/web/html")
	// 使用runtime获取main.go路径，进而获得绝对路径，但打包后获取为当前文件路径
	//if _, filePath, _, ok := runtime.Caller(0); ok {
	//	curDir := path.Dir(filePath)
	//	log.Logf("filePath%v", filePath)
	//	log.Logf("curDir%v", curDir)
	//	service.Handle("/", http.FileServer(http.Dir(curDir+"/html")))
	//} else {
	//	log.Log("html dir err:main.go file path nil")
	//}
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
