# Micro [github.com/micro](http://github.com/micro)
[![Slack](https://img.shields.io/badge/slack-join-D60051.svg)](https://hbchen.slack.com/messages/CE68CJ60Z)\

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
cd service/account

go run main.go --profile starter-local
```

Test example service 
```shell script
curl "http://localhost:8080/account/info?id=1"
```

## Kubernetes

> TODO