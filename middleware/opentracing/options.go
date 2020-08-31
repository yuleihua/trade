package opentracing

type Option func(c *Options)

type Options struct {
	ServiceName string
	Address     string
}

func ServiceName(serviceName string) Option {
	return func(c *Options) {
		c.ServiceName = serviceName
	}
}

func Address(address string) Option {
	return func(c *Options) {
		c.Address = address
	}
}

func applyOptions(otType TracerType, options ...Option) Options {
	opts := Options{}
	for _, option := range options {
		option(&opts)
	}

	if len(opts.ServiceName) == 0 {
		opts.ServiceName = "github.com/yuleihua/trade"
	}

	// jaeger-agent 127.0.0.1:6831
	if len(opts.Address) == 0 {
		switch otType {
		case TracerJaeger:
			opts.Address = "127.0.0.1:6831"
		}
	}

	return opts
}
