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
# Account
cd service/account
go run main.go --profile starter-local

# Greeting
cd service/greeting
go run main.go --profile starter-local
```

Test example service 
```shell script
# Account
curl "http://localhost:8080/account/info?id=1"

# Greeting
curl "http://localhost:8080/greeting/call?msg=helloworld"
{"id":"1","msg":"helloworld"}
```

## Kubernetes

> Attention: default ingress class=nginx, host=api.micro.hbchen.com

```shell
helm install -n micro micro-server cicd/charts/micro \
--set ingress.enabled=true

# Digest
helm install -n micro micro-server cicd/charts/micro \
--set image.tag="latest@sha256:1e2c8df50398c2dcd4b96065b8b81842a86c5cc83d8ef1ae96ac7b5d8432add3" \
--set ingress.enabled=true
```

```shell
helm install -n micro micro-example cicd/charts/servic

# Digest
helm install -n micro micro-example cicd/charts/service \
--set image.tag="latest@sha256:fa2e56f01a4704ad298331cd0356a0e174e9358701922552934bbd4987c9fb80"
```

```shell
curl "http://api.micro.hbchen.com/account/info?id=1"
```