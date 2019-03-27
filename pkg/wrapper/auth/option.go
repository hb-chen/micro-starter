package auth

import ()

type Options struct {
	serviceName string
	skipper     SkipperFunc
	tokenLookup TokenLookupFunc
}

type Option func(*Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		serviceName: "go.micro.srv.auth",
		skipper:     DefaultSkipperFunc,
		tokenLookup: DefaultTokenLookupFunc,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func ServiceName(sn string) Option {
	return func(o *Options) {
		o.serviceName = sn
	}
}

func Skipper(skipper SkipperFunc) Option {
	return func(o *Options) {
		o.skipper = skipper
	}
}

func TokenLookup(tf TokenLookupFunc) Option {
	return func(o *Options) {
		o.tokenLookup = tf
	}
}
