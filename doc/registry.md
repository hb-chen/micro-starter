### Registry

##### Interface
```go
// The registry provides an interface for service discovery
// and an abstraction over varying implementations
// {consul, etcd, zookeeper, ...}
type Registry interface {
	Register(*Service, ...RegisterOption) error
	Deregister(*Service) error
	GetService(string) ([]*Service, error)
	ListServices() ([]*Service, error)
	Watch() (Watcher, error)
	String() string
}
```

##### Options
```go
type Options struct {
	Addrs     []string
	Timeout   time.Duration
	Secure    bool
	TLSConfig *tls.Config

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

type RegisterOptions struct {
	TTL time.Duration
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}
```