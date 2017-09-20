# micro
Go Micro 应用服务化治理实践

### Protocol
```bash
go get github.com/micro/protobuf/{proto,protoc-gen-go}
protoc --go_out=plugins=micro:. greeter.proto
```

micro new demo

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
	
### 模块创建
micro new --type srv --alias account github.com/hb-go/micro/account/srv
micro new --type api --alias account github.com/hb-go/micro/account/api
micro new --type web --alias account github.com/hb-go/micro/account/web