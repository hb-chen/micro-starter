
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

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

.PHONY: micro
micro:
	./dist/micro_$(GOOS)_$(GOARCH)/bin/micro --profile starter-local server

.PHONY: example
example:
	./dist/example_$(GOOS)_$(GOARCH)/bin/example --profile starter-local

.PHONY: release
release:
	goreleaser release --config $(goreleaser_config) --skip-validate --skip-publish --rm-dist

.PHONY: snapshot
snapshot:
	goreleaser release --config $(goreleaser_config) --skip-publish --snapshot --rm-dist

.PHONY: test
test:
	go test -race -cover -v ./cmd/... ./service/...  ./profile/... ./pkg/...

.PHONY: lint
lint:
	golangci-lint run ./cmd/... ./service/...  ./profile/... ./pkg/...

.PHONY: pack_build
pack_build:
	pack build micro \
	--builder paketobuildpacks/builder:tiny \
	--descriptor manifests/buildpacks/project.toml \
	--tag registry.cn-hangzhou.aliyuncs.com/hb-chen/micro-starter-micro:latest

.PHONY: run
run:
	docker run micro --profile starter-local server