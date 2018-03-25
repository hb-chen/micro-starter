# micro
Go Micro 应用服务化治理实践

[github.com/micro/micro](http://github.com/micro/micro)
[github.com/micro/go-micro](http://github.com/micro/go-micro)

##### [go-micro](/doc/README.md)
![go-micro](/doc/img/micro.jpg "go-micro")

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
micro new --type srv --alias account github.com/hb-go/micro/auth/srv
micro new --type api --alias account github.com/hb-go/micro/account/api
micro new --type web --alias account github.com/hb-go/micro/account/web

micro new --type srv --alias account github.com/hb-go/micro/post/srv
micro new --type api --alias account github.com/hb-go/micro/post/api
micro new --type web --alias account github.com/hb-go/micro/post/web
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
- [Kafka&ZooKeeper安装使用](https://kafka.apache.org/quickstart)
```bash
# create topic
$ bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic zipkin
```
- [Zipkin](https://github.com/openzipkin/zipkin)
- [Kafka Collector](https://github.com/openzipkin/zipkin/tree/master/zipkin-server#kafka-collector)
```bash
# Docker 运行zipkin
# 需要指定KAFKA_ZOOKEEPER，host使用主机IP
$ docker run --name zipkin -d -p 9411:9411 \
--env KAFKA_ZOOKEEPER=192.168.1.1:2181 \
openzipkin/zipkin

#localhost:9411 查看Trace信息
```