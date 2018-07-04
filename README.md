# Micro [github.com/micro](http://github.com/micro)
Go Micro 应用服务化治理实践

<a href="/doc/README.md">![go-micro](/doc/img/micro.jpg "go-micro")</a>

## 运行示例

### 服务发现-Consul
```bash
$ consul agent -dev -advertise 127.0.0.1
```

### 运行
```bash
$ micro api
$ micro web

# Auth SRV
$ cd auth/srv/ && go run main.go

# Account API
$ cd account/api/ && go run main.go
$ curl -H 'Content-Type: application/json' \
            -H "Authorization: Bearer VALID_TOKEN" \
            -d '{"nickname": "Hobo", "pwd": "pwd"}' \
             http://localhost:8080/login
             
# Account Web
$ cd account/web/ && go run main.go
http://localhost:8082/accout

# Post SRV
$ cd post/srv/ && go run main.go

# Post API
$ cd post/api/ && go run main.go
http://localhost:8080/post?id=1
http://localhost:8080/post/comments?id=1

# Post Web
$ cd post/web/ && go run main.go
http://localhost:8082/post
```

### Echo做Web框架
- [Micro＋Echo示范](/_echo-web)
- 更多Echo示例☞[hb-go/echo-web](https://github.com/hb-go/echo-web)

### 插件替换
> 需要替换的插件import到plugins.go

> micro工具需要手动编译
```bash
# 编译micro
$ go build -i -o micro ./main.go ./plugins.go

# 运行micro api/web
$ micro --transport=tcp api
$ micro --transport=tcp web

# 运行go-micro服务
$ go run main.go plugins.go --transport=tcp
```
```go
// plugins.go

package main

import (
	// registry
	// k8s
	_ "github.com/micro/go-plugins/registry/kubernetes"
	// etcd v3
	//_ "github.com/micro/go-plugins/registry/etcdv3"

	// transport
	// tcp
	_ "github.com/micro/go-plugins/transport/tcp"
	// nats
	//_ "github.com/micro/go-plugins/transport/nats"

	// broker
	// kafka
	//_ "github.com/micro/go-plugins/broker/kafka"
)
```

## 框架使用

### 模块创建
```bash
$ micro new --type srv --alias auth github.com/hb-go/micro/auth/srv
$ micro new --type api --alias account github.com/hb-go/micro/account/api
$ micro new --type web --alias account github.com/hb-go/micro/account/web

$ micro new --type srv --alias post github.com/hb-go/micro/post/srv
$ micro new --type api --alias post github.com/hb-go/micro/post/api
$ micro new --type web --alias post github.com/hb-go/micro/post/web
```

### Protobuf [protoc-gen-micro](https://github.com/micro/protoc-gen-micro)
```bash
$ go get github.com/micro/protoc-gen-micro

$ protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. post/srv/proto/example/example.proto
```

##### .proto批量处理
```bash
# 批处理工具打包
$ go build -i -o build/bin/proto_batch tools/proto/batch.go

# ./build/bin/proto_batch -h
$ ./build/bin/proto_batch -r auth:account:post
```

### API
```bash
$ micro api
$ micro --enable_stats api
    http://localhost:8080/stats
$ micro api --namespace=com.example.api

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
```

### Web
```bash
$ micro --enable_stats web
```
	
### Trace
[Jaeger](http://jaeger.readthedocs.io/en/latest/getting_started/#all-in-one-docker-image)
```bash
$ docker run -d --name=jaeger -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p5775:5775/udp -p6831:6831/udp -p6832:6832/udp   -p5778:5778 -p16686:16686 -p14268:14268 -p9411:9411 jaegertracing/all-in-one:latest
#http://localhost:16686 查看Trace信息
```