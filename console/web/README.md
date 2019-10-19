# Example Service

This is the Example service

Generated with

```
micro new github.com/hb-go/micro/console/web --namespace=go.micro --alias=example --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.web.example
- Type: web
- Alias: example

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./example-web
```

Build a docker image
```
make docker
```