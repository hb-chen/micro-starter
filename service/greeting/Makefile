
GOPATH:=$(shell go env GOPATH)

ifeq ($(OS),Windows_NT)
    uname_S := Windows
else
    uname_S := $(shell uname -s)
endif

# Goreleaser config
ifeq ($(uname_S), Darwin)
    goreleaser_config = .goreleaser.yml
else
	goreleaser_config = .goreleaser.yml
endif

.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v3/cmd/protoc-gen-micro
	go get github.com/micro/micro/v3/cmd/protoc-gen-openapi

.PHONY: api
api:
	protoc --proto_path=${GOPATH}/src/ --openapi_out=. github.com/hb-chen/micro-starter/service/greeting/proto/greeting/greeting.proto

.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src/ --micro_out=. --go_out=. github.com/hb-chen/micro-starter/service/greeting/proto/greeting/greeting.proto

.PHONY: build
build: proto
	go build -o bin/console  main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t example-srv:latest

.PHONY: release
release:
	goreleaser release --config $(goreleaser_config) --skip-validate --skip-publish --rm-dist

.PHONY: snapshot
snapshot:
	goreleaser release --config $(goreleaser_config) --skip-publish --snapshot --rm-dist