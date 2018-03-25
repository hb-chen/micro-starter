### Client

##### Interface
```go
// Client is the interface used to make requests to services.
// It supports Request/Response via Transport and Publishing via the Broker.
// It also supports bidiectional streaming of requests.
type Client interface {
	Init(...Option) error
	Options() Options
	NewPublication(topic string, msg interface{}) Publication
	NewRequest(service, method string, req interface{}, reqOpts ...RequestOption) Request
	NewProtoRequest(service, method string, req interface{}, reqOpts ...RequestOption) Request
	NewJsonRequest(service, method string, req interface{}, reqOpts ...RequestOption) Request
	Call(ctx context.Context, req Request, rsp interface{}, opts ...CallOption) error
	CallRemote(ctx context.Context, addr string, req Request, rsp interface{}, opts ...CallOption) error
	Stream(ctx context.Context, req Request, opts ...CallOption) (Streamer, error)
	StreamRemote(ctx context.Context, addr string, req Request, opts ...CallOption) (Streamer, error)
	Publish(ctx context.Context, p Publication, opts ...PublishOption) error
	String() string
}

// Publication is the interface for a message published asynchronously
type Publication interface {
	Topic() string
	Message() interface{}
	ContentType() string
}

// Request is the interface for a synchronous request used by Call or Stream
type Request interface {
	Service() string
	Method() string
	ContentType() string
	Request() interface{}
	// indicates whether the request will be a streaming one rather than unary
	Stream() bool
}

// Streamer is the inteface for a bidirectional synchronous stream
type Streamer interface {
	Context() context.Context
	Request() Request
	Send(interface{}) error
	Recv(interface{}) error
	Error() error
	Close() error
}
```

##### Options
```go
type Options struct {
	// Used to select codec
	ContentType string

	// Plugged interfaces
	Broker    broker.Broker
	Codecs    map[string]codec.NewCodec
	Registry  registry.Registry
	Selector  selector.Selector
	Transport transport.Transport

	// Connection Pool
	PoolSize int
	PoolTTL  time.Duration

	// Middleware for client
	Wrappers []Wrapper

	// Default Call Options
	CallOptions CallOptions

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

type CallOptions struct {
	SelectOptions []selector.SelectOption

	// Backoff func
	Backoff BackoffFunc
	// Check if retriable func
	Retry RetryFunc
	// Transport Dial Timeout
	DialTimeout time.Duration
	// Number of Call attempts
	Retries int
	// Request/Response timeout
	RequestTimeout time.Duration

	// Middleware for low level call func
	CallWrappers []CallWrapper

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

type PublishOptions struct {
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

type RequestOptions struct {
	Stream bool

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}
```