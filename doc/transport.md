### Transport

##### Interface
```go
type Message struct {
	Header map[string]string
	Body   []byte
}

type Socket interface {
	Recv(*Message) error
	Send(*Message) error
	Close() error
}

type Client interface {
	Socket
}

type Listener interface {
	Addr() string
	Close() error
	Accept(func(Socket)) error
}

// Transport is an interface which is used for communication between
// services. It uses socket send/recv semantics and had various
// implementations {HTTP, RabbitMQ, NATS, ...}
type Transport interface {
	Dial(addr string, opts ...DialOption) (Client, error)
	Listen(addr string, opts ...ListenOption) (Listener, error)
	String() string
}
```

##### Options
```go
type Options struct {
	Addrs     []string
	Codec     codec.Codec
	Secure    bool
	TLSConfig *tls.Config
	// Timeout sets the timeout for Send/Recv
	Timeout time.Duration
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

type DialOptions struct {
	Stream  bool
	Timeout time.Duration

	// TODO: add tls options when dialling
	// Currently set in global options

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

type ListenOptions struct {
	// TODO: add tls options when listening
	// Currently set in global options

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}
```