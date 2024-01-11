# Micro [github.com/micro](http://github.com/micro)

[![Slack](https://img.shields.io/badge/slack-join-D60051.svg)](https://hbchen.slack.com/messages/CE68CJ60Z)

![go-micro](/doc/img/micro.jpg "go-micro")

## Local

Build micro cmd
```shell script
go build -o bin/micro cmd/micro/main.go
```

Start registry & api with server runtime
```shell script
./bin/micro --profile starter-local server
```

<details>
  <summary> Start registry & api with service command </summary>
Run registry service
```shell script
./bin/micro --profile starter-local service registry
```

Run API service
```shell script
./bin/micro --profile starter-local service api
```
</details>

Run example service
```shell script
# Greeting
cd service/greeting
CGO_ENABLED=0 go run main.go --profile starter-local
```

Test example service
```shell script
# Greeting
curl "http://localhost:8080/greeting/call?msg=helloworld"
{"id":"1","msg":"helloworld"}

curl "http://localhost:8080/greeting/list?page=1&size=10"
{"items":[{"id":"1","msg":"helloworld"}]}
```

## Kubernetes

```shell
$ make snapshot

$ docker build ./ -f Dockerfile --platform=linux/amd64 -t registry.cn-hangzhou.aliyuncs.com/hb-chen/micro-starter-micro:latest
```

> Attention: default ingress class=nginx, host=api.micro.hbchen.com

```shell
helm install -n micro micro-server manifests/charts/micro \
--set ingress.enabled=true

# Digest
helm install -n micro micro-server manifests/charts/micro \
--set image.tag="latest@sha256:aceabd67ac333dcd19bde3524c54e7a556b8651cf049495ab6e086d45bb7ad77" \
--set ingress.enabled=true
```

```shell
helm install -n micro micro-example manifests/charts/service

# Digest
helm install -n micro micro-example manifests/charts/service \
--set image.tag="latest@sha256:a2af30ff9a0a66ade77672e01679a2b02ead3b2b0f27bd7092d726d75fd069e0"
```

```shell
curl "http://api.micro.hbchen.com/greeting/call?msg=helloworld"
```