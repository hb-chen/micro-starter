# Greeting Service

This is the Greeting service

## Usage

Generate the proto code

```
make proto
```

Run the service

```
go run main.go --profile starter-local
```

Call service

```shell script
curl "http://localhost:8080/greeting/call?msg=helloworld"
{"id":"1","msg":"helloworld"}
```

## Image

```shell
$ make snapshot

$ docker build ./ -f Dockerfile  --platform=linux/amd64 -t registry.cn-hangzhou.aliyuncs.com/hb-chen/micro-starter-example:latest
```

