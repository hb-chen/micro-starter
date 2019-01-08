## go-micro服务加入service mesh示范

使用go-micro的`client`&`server`istio插件: [hb-go/micro-plugins](https://github.com/hb-go/micro-plugins)

### TODO
- micro-plugins
    - FQDN Server/Service address定义

## K8s
```bash
# http/gRPC部署
$ kubectl apply -f service_deployment.yaml
$ kubectl apply -f destination-rule.yaml
$ kubectl apply -f virtual-service.yaml
$ kubectl apply -f gateway.yaml

# http验证
$ curl -H "Content-Type:application/json" -X GET http://192.168.99.100:31380/example/call?name=Hobo
{"statusCode":200,"body":"{\"msg\":\"Hello Hobo\"}"}

$ curl -H "Content-Type:application/json" -X POST -d '{"name":"Hobo"}' http://192.168.99.100:31380/example/call
{"statusCode":200,"body":"{\"msg\":\"Hello Hobo\"}"}

# gRPC验证
$ go run grpc_client.go -address 192.168.99.100:31380
2019/01/08 23:12:09 resp: statusCode:200 body:"{\"msg\":\"Hello Hobo\"}" 
2019/01/08 23:12:09 duration: 17.664807ms
```

## Local测试

### HTTP
使用[sofa-mosn](https://github.com/alipay/sofa-mosn)做代理调试，参考[sofa-mosn/examples/http-sample](https://github.com/alipay/sofa-mosn/tree/master/examples/http-sample)
- 需要一个编译好的SOFAMosn程序`sofa-mosn`
```bash
$ cd pwd/http

# run api service
$ cd api
$ go run main.go -server_address localhost:8081 -client_call_address localhost:2046

# run srv service
$ cd srv
$ go run main.go -server_address localhost:8082

# run api mosn 代理
$ ./sofa-mosn start -c mosn_api.json

# run srv mosn 代理
$ ./sofa-mosn start -c mosn_srv.json
```

#### 测试
```bash
$ curl -H "Content-Type:application/json" -X GET http://127.0.0.1:2045/example/call?name=Hobo
{"statusCode":200,"body":"{\"msg\":\"Hello Hobo\"}"}

$ curl -H "Content-Type:application/json" -X POST -d '{"name":"Hobo"}' http://127.0.0.1:2045/example/call
  {"statusCode":200,"body":"{\"msg\":\"Hello Hobo\"}"}
```

### gRPC
> TODO
- 使用Envoy做代理调试

```bash
$ cd pwd/grpc

# run api service
$ cd api
$ go run main.go -server_address localhost:8082 -client_call_address localhost:8081

# run srv service
$ cd srv
$ go run main.go -server_address localhost:8081
```

#### 测试
```bash
$ go run grpc_client.go -address localhost:8082
2019/01/08 23:12:09 resp: statusCode:200 body:"{\"msg\":\"Hello Hobo\"}" 
2019/01/08 23:12:09 duration: 17.664807ms
```
