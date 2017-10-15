# micro
Go Micro 应用服务化治理实践

[github.com/micro/micro](github.com/micro/micro)

### 模块创建
```bash
micro new --type srv --alias account github.com/hb-go/micro/account/srv
micro new --type api --alias account github.com/hb-go/micro/account/api
micro new --type web --alias account github.com/hb-go/micro/account/web
```

### Protobuf [GRPC Gateway](https://micro.mu/docs/grpc-gateway.html)
```bash
go get github.com/micro/protobuf/{proto,protoc-gen-go}
protoc --go_out=plugins=micro:. account/srv/proto/example/example.proto

# api中import "github.com/micro/go-api/proto/api.proto";
# 报错:github.com/micro/go-api/proto/api.proto: File not found.
# 需要增加依赖的路径 -I$GOPATH/src \
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  --go_out=plugins=micro:. \
  post/api/proto/example/example.proto
```

### 服务发现-Consul
```
consul agent -dev -advertise 127.0.0.1
```

### Running
```bash
# Post API
go run post/srv/main.go
go run post/api/main.go
micro api
http://localhost:8080/post/post/post?id=1

# Post Web
go run post/web/main.go
micro web
http://localhost:8082/post
```

### API
	micro api
	micro --enable_stats api
		http://localhost:8080/stats
	micro api --namespace=com.example.api

		Make a HTTP call
		curl "http://localhost:8080/greeter/say/hello?name=Asim+Aslam"

		Make an RPC call via the /rpc
		curl -d 'service=go.micro.srv.greeter' \
			-d 'method=Say.Hello' \
			-d 'request={"name": "Asim Aslam"}' \
			http://localhost:8080/rpc

		Make an RPC call via /rpc with content-type set to json
		$ curl -H 'Content-Type: application/json' \
			-d '{"service": "go.micro.srv.greeter", "method": "Say.Hello", "request": {"name": "Asim Aslam"}}' \
			http://localhost:8080/rpc

### Web
	micro --enable_stats web