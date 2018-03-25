### Server

##### Interface
```go
type Server interface {
	Options() Options
	Init(...Option) error
	Handle(Handler) error
	NewHandler(interface{}, ...HandlerOption) Handler
	NewSubscriber(string, interface{}, ...SubscriberOption) Subscriber
	Subscribe(Subscriber) error
	Register() error
	Deregister() error
	Start() error
	Stop() error
	String() string
}

type Publication interface {
	Topic() string
	Message() interface{}
	ContentType() string
}

type Request interface {
	Service() string
	Method() string
	ContentType() string
	Request() interface{}
	// indicates whether the request will be streamed
	Stream() bool
}

// Streamer represents a stream established with a client.
// A stream can be bidirectional which is indicated by the request.
// The last error will be left in Error().
// EOF indicated end of the stream.
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
	Codecs       map[string]codec.NewCodec
	Broker       broker.Broker
	Registry     registry.Registry
	Transport    transport.Transport
	Metadata     map[string]string
	Name         string
	Address      string
	Advertise    string
	Id           string
	Version      string
	HdlrWrappers []HandlerWrapper
	SubWrappers  []SubscriberWrapper

	RegisterTTL time.Duration

	// Debug Handler which can be set by a user
	DebugHandler debug.DebugHandler

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}
```