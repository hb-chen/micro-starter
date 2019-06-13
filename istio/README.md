## go-micro服务加入service mesh示范

使用go-micro的`client`&`server`istio插件: [hb-go/micro-plugins](https://github.com/hb-go/micro-plugins)

### TODO
- gRPC版本升级
- micro-plugins
    - FQDN Server/Service address定义

## Local测试
使用以下流程模拟服务在`mesh`环境的运行
```bash
[|envoy:2045| -> |api:8081| -> |envoy:2046|] -> [|envoy:2047| -> |srv:8082|]
```

### HTTP

参考`Envoy`官方文档[Using The Envoy Docker Image](https://www.envoyproxy.io/docs/envoy/v1.10.0/start/start#using-the-envoy-docker-image)，熟悉如何在Docker环境使用。

#### Envoy代理
```bash
$ cd {pwd}/envoy/http
```
在`envoy.yaml`配置中`8081`和`8082`两个服务是两个测试用例`api`和`srv`服务，所以需要根据自身环境修改`address`，这里测试在主机上运行服务，所以配置为主机IP`192.168.1.110`。

```bash
# Docker镜像
$ docker build -t envoy:v1 .

# 运行Envoy
$ docker run --name envoy -p 9901:9901 -p 2045:2045 -p 2046:2046 envoy:v1
```

***~~MOSN代理~~***

使用[sofa-mosn](https://github.com/alipay/sofa-mosn)做代理调试，参考[sofa-mosn/examples/http-sample](https://github.com/alipay/sofa-mosn/tree/master/examples/http-sample)
- 需要一个编译好的SOFAMosn程序`sofa-mosn`
```bash
# run api mosn 代理
$ ./sofa-mosn start -c mosn/http/mosn_api.json

# run srv mosn 代理
$ ./sofa-mosn start -c mosn/http/mosn_srv.json
```

#### 运行服务
```bash
$ cd {pwd}/http

# run api service
$ cd api
$ go run main.go -server_address localhost:8081 -client_call_address localhost:2046

# run srv service
$ cd srv
$ go run main.go -server_address localhost:8082
```

#### 测试
```bash
$ curl -H "Content-Type:application/json" -X GET http://127.0.0.1:2045/example/call?name=Hobo
{"statusCode":200,"body":"{\"msg\":\"Hello Hobo\"}"}

$ curl -H "Content-Type:application/json" -X POST -d '{"name":"Hobo"}' http://127.0.0.1:2045/example/call
  {"statusCode":200,"body":"{\"msg\":\"Hello Hobo\"}"}
```

### gRPC
> 由于`gRPC`版本升级，当前`istio-grpc`插件版本不可用，`k8s`用例为旧的镜像，应该可以测试

#### 运行服务
```bash
$ cd {pwd}/grpc

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

## K8s
> `k8s`用例为旧的镜像

```bash
# http/gRPC部署
$ kubectl apply -f service-deployment.yaml
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

# HTTP Gateway开启JWT
$ kubectl apply -f gateway-jwt.yaml
$ TOKEN=$(curl https://raw.githubusercontent.com/istio/istio/release-1.1/security/tools/jwt/samples/demo.jwt -s)
$ curl --header "Authorization: Bearer $TOKEN" -H "Content-Type:application/json" -X GET http://192.168.99.100:31380/example/call?name=Hobo
{"statusCode":200,"body":"{\"msg\":\"Hello Hobo\"}"}
$ curl --header "Authorization: Bearer $TOKEN" -H "Content-Type:application/json" -X POST -d '{"name":"Hobo"}' http://192.168.99.100:31380/example/call
{"statusCode":200,"body":"{\"msg\":\"Hello Hobo\"}"}
```