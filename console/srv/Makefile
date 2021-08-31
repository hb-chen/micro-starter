
GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v3/cmd/protoc-gen-micro
	go get github.com/micro/micro/v3/cmd/protoc-gen-openapi

.PHONY: api
api:
	protoc --proto_path=${GOPATH}/src/ --openapi_out=. github.com/hb-chen/micro-starter/console/srv/proto/account/account.proto

.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src/ --micro_out=. --go_out=. github.com/hb-chen/micro-starter/console/srv/proto/account/account.proto

.PHONY: build
build: proto

	go build -o bin/console  main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t example-srv:latest
