package gosocket

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/gorilla/websocket"
)

// Namespace is the key to use to store and access the custom config data
const Namespace = "github.com/troy-johnson/krakend-websocket"

// Config is the custom config struct containing the params for the sony/gobreaker package
type Config struct {
	Interval        int
	Timeout         int
	MaxErrors       int
	LogStatusChange bool
}

// ZeroCfg is the zero value for the Config struct
var ZeroCfg = Config{}

// ConfigGetter implements the config.ConfigGetter interface. It parses the extra config for the
// gobreaker adapter and returns a ZeroCfg if something goes wrong.
func ConfigGetter(e config.ExtraConfig) interface{} {
	v, ok := e[Namespace]
	if !ok {
		return ZeroCfg
	}
	tmp, ok := v.(map[string]interface{})
	if !ok {
		return ZeroCfg
	}
	cfg := Config{}
	if v, ok := tmp["interval"]; ok {
		switch i := v.(type) {
		case int:
			cfg.Interval = i
		case float64:
			cfg.Interval = int(i)
		}
	}
	if v, ok := tmp["timeout"]; ok {
		switch i := v.(type) {
		case int:
			cfg.Timeout = i
		case float64:
			cfg.Timeout = int(i)
		}
	}
	if v, ok := tmp["maxErrors"]; ok {
		switch i := v.(type) {
		case int:
			cfg.MaxErrors = i
		case float64:
			cfg.MaxErrors = int(i)
		}
	}
	value, ok := tmp["logStatusChange"].(bool)
	cfg.LogStatusChange = ok && value

	return cfg
}

type WebSocket struct {
}

// NewCircuitBreaker builds a gobreaker circuit breaker with the injected config
func NewWebSocket(logger logging.Logger) *socket.WebSocket {
	// settings := gosocket.Settings{
	// 	Interval: time.Duration(cfg.Interval) * time.Second,
	// 	Timeout:  time.Duration(cfg.Timeout) * time.Second,
	// 	ReadyToTrip: func(counts gobreaker.Counts) bool {
	// 		return counts.ConsecutiveFailures > uint32(cfg.MaxErrors)
	// 	},
	// }

	// if cfg.LogStatusChange {
	// 	settings.OnStateChange = func(name string, from gobreaker.State, to gobreaker.State) {
	// 		logger.Warning(fmt.Sprintf("circuit breaker named '%s' went from '%s' to '%s'", name, from.String(), to.String()))
	// 	}
	// }

	return WebSocket
}
