package proxy

import (
	"context"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"

	gcb "github.com/troy-johnson/krakend-websocket"
)

// BackendFactory adds a cb middleware wrapping the internal factory
func BackendFactory(next proxy.BackendFactory, logger logging.Logger) proxy.BackendFactory {
	return func(cfg *config.Backend) proxy.Proxy {
		return NewMiddleware(cfg, logger)(next(cfg))
	}
}

// NewMiddleware builds a middleware based on the extra config params or fallbacks to the next proxy
func NewMiddleware(remote *config.Backend, logger logging.Logger) proxy.Middleware {
	// data := gcb.ConfigGetter(remote.ExtraConfig).(gcb.Config)
	// if data == gcb.ZeroCfg {
	// 	return proxy.EmptyMiddleware
	// }
	cb := gcb.WebSocket(logger)

	return func(next ...proxy.Proxy) proxy.Proxy {
		if len(next) > 1 {
			panic(proxy.ErrTooManyProxies)
		}
		return func(ctx context.Context, request *proxy.Request) (*proxy.Response, error) {
			result, err := cb.Execute(func() (interface{}, error) { return next[0](ctx, request) })
			if err != nil {
				return nil, err
			}
			return result.(*proxy.Response), err
		}
	}
}
