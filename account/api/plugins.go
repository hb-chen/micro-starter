package main

import (
	// registry
	// k8s
	//_ "github.com/micro/go-plugins/registry/kubernetes"
	// etcd v3
	_ "github.com/micro/go-plugins/registry/etcd"

	// transport
	// tcp
	_ "github.com/micro/go-plugins/transport/tcp"
	// nats
	//_ "github.com/micro/go-plugins/transport/nats"

	// broker
	// kafka
	//_ "github.com/micro/go-plugins/broker/kafka"
)
