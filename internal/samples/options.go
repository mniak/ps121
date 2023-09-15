package samples

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mniak/duplicomp/internal/noop"
	"github.com/mniak/duplicomp/internal/samples/internal"
)

type _ServerHandler func(ctx context.Context, ping *internal.Ping) (*internal.Pong, error)

type _Options struct {
	Port                 int
	Logger               *log.Logger
	ServerHandlerFactory func(*log.Logger) _ServerHandler
}
type _Option func(*_Options)

func WithHandlerFactory(hf func(*log.Logger) _ServerHandler) _Option {
	return func(o *_Options) {
		o.ServerHandlerFactory = hf
	}
}

func WithPort(port int) _Option {
	return func(o *_Options) {
		o.Port = port
	}
}

func WithName(name string) _Option {
	return func(o *_Options) {
		o.Logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", name), 0)
	}
}

func (o *_Options) apply(opts ..._Option) *_Options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func buildOptions(opts ..._Option) _Options {
	var options _Options
	options.apply(opts...)
	return options
}

func defaultOptions() *_Options {
	return &_Options{
		Port:                 9000,
		Logger:               noop.Logger(),
		ServerHandlerFactory: defaultServerHandler,
	}
}
