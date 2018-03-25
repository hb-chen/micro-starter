### Broker

##### Interface
```go
// Broker is an interface used for asynchronous messaging.
// Its an abstraction over various message brokers
// {NATS, RabbitMQ, Kafka, ...}
type Broker interface {
	Options() Options
	Address() string
	Connect() error
	Disconnect() error
	Init(...Option) error
	Publish(string, *Message, ...PublishOption) error
	Subscribe(string, Handler, ...SubscribeOption) (Subscriber, error)
	String() string
}

// Handler is used to process messages via a subscription of a topic.
// The handler is passed a publication interface which contains the
// message and optional Ack method to acknowledge receipt of the message.
type Handler func(Publication) error

type Message struct {
	Header map[string]string
	Body   []byte
}

// Publication is given to a subscription handler for processing
type Publication interface {
	Topic() string
	Message() *Message
	Ack() error
}

// Subscriber is a convenience return type for the Subscribe method
type Subscriber interface {
	Options() SubscribeOptions
	Topic() string
	Unsubscribe() error
}
```

##### Options
```go
type Options struct {
	Addrs     []string
	Secure    bool
	Codec     codec.Codec
	TLSConfig *tls.Config
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}
```