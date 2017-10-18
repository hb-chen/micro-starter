#!/bin/bash

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o account-web ./main.go
docker build -t hbchen/micro-account-web:v0.0.4 .
docker push hbchen/micro-account-web