
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:

	go build -o example-web main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t example-web:latest

.PHONY: vue
vue:
	cd vue && npm run build:prod

.PHONY: statik
statik:
	go get github.com/rakyll/statik
	go generate ./...
