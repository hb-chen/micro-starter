# 框架介绍

## [micro](https://github.com/micro/micro)
主要由以下工具构成:
- API
    > API网关提供HTTP服务，并将请求路由到指定的微服务，是Micro的统一入口，可以用作反向代理，或者将HTTP请求转到RPC

- Web
    > 为Micro提供一套仪表盘，并可作为Web应用的反向代理，有别于普通API将Web应用作为了Micro的一等公民，其实和API差不多，但同时提供了对socket的支持

- Sidecar
    > 通过HTTP的方式实现go-micro全部功能的RPC代理，通过Sidecar可以方便的将其他语言集成到Micro框架中，方便解决解决应用的异构框架问题

- CLI
    > 提供一套命令行工具，可以方便的与Micro服务进行交互，并且可以通过Sidecar代理CLI命令

- Bot
    > 通过Bot可以在Micro环境中方便的与Slack、HipChat、XMPP等进行集成，通过消息的方式模仿CLI功能 

## [go-micro](https://github.com/micro/go-micro)
go-micro是Micro的核心，是一套Go语言的可插拔RPC框架，提供服务发现、负载均衡、同步/异步通信、编码、服务接口等，所有组件均设计为Interface，便于扩展

![go-micro](/doc/img/micro.jpg)

主要有以下组件构成:
- [Registry](/doc/registry.md)
    > 提供一套服务注册、发现、注销、监测机制，服务注册中心支持consul、etcd2/3、zookeeper、gossip、k8s、eureka等

- [Selector](/doc/selector.md)
    > 选择器提供了负载均衡，可以通过过滤方法对微服务进行过滤，并通过不同路由算法选择微服务，以及缓存等

- [Transport](/doc/transport.md)
    > 微服务间同步请求/响应通信方式，相对Go标准net包做了更高的抽象，支持更多的传输方式，如http、grpc、tcp、udp、Rabbitmq等

- [Broker](/doc/broker.md)
    > 微服务键异步发布/订阅通信方式，更好的处理分布式系统解耦问题，默认使用http方式，生产环境通常会使用消息中间件，如Kafka、RabbitMQ、NSQ等

- Codec
    > 服务间消息的编解码，支持json、protobuf、bson、msgpack等，与普通编码格式不同都是支持RPC格式

- [Server](/doc/server.md)
    > 用于启动服务，为服务命名、注册Handler、添加中间件等

- [Client](/doc/client.md)
    > 提供微服务客户端，通过Registry、Selector、Transport、Broker实现以服务名来查找服务、负载均衡、同步通信、异步消息等

### [go-plugins](https://github.com/micro/go-plugins)
go-plugins是Micro的插件库，除go-micro相应组件的扩展外，还有其他如Trace、KV存储、监控等

### [ja-micro](https://github.com/Sixt/ja-micro)
Micro生态RPC框架的Java实现，使Java框架与Go一样方便的集成到Micro中

### [go-api](https://github.com/micro/go-api)

### [go-os](https://github.com/micro/go-os)

### [go-grpc](https://github.com/micro/go-grpc)