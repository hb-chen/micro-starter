## go-micro服务加入service mesh示范

使用sofa-mosn做代理演示，参考[sofa-mosn/examples/http-sample](https://github.com/alipay/sofa-mosn/tree/master/examples/http-sample)

### 准备

- 需要一个编译好的SOFAMosn程序`sofa-mosn`

### 运行示例
```bash
$ cd pwd

# run api service
$ cd api
$ go run main.go

# run srv service
$ cd api
$ go run main.go

# run api mosn 代理
$ ./sofa-mosn start -c mosn_api.json

# run srv mosn 代理
$ ./sofa-mosn start -c mosn_srv.json
```

### 测试
```bash
$ curl http://127.0.0.1:2045/Example.Call
{"statusCode":200,"body":"{\"msg\":\"Hello \"}"}

# TODO go-api支持
$ curl -H "Content-Type:application/json" -X POST -d '{"post":{ "name":{"key":"name","values":["Hobo"]}}}' http://127.0.0.1:2045/Example.Call
{"statusCode":200,"body":"{\"msg\":\"Hello Hobo\"}"}
```
