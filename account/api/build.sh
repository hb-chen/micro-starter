#!/bin/bash

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o account-api ./main.go
docker build -t hbchen/micro-account-api:v0.0.1 .
docker push hbchen/micro-account-api