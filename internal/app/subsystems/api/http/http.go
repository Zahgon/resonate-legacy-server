package http

import (
	"net"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-viper/mapstructure/v2"
	"github.com/resonatehq/resonate/internal/app/subsystems/api"
	"github.com/resonatehq/resonate/internal/kernel/t_api"
	"github.com/resonatehq/resonate/internal/metrics"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	i_api "github.com/resonatehq/resonate/internal/api"
)

type Config struct {
	Addr          string            `flag:"addr" desc:"http server address" default:":8001"`
	Auth          map[string]string `flag:"auth" desc:"http basic auth username password pairs"`
	Cors          Cors              `flag:"cors" desc:"http cors settings"`
	Timeout       time.Duration     `flag:"timeout" desc:"http server graceful shutdown timeout" default:"10s"`
	TaskFrequency time.Duration     `flag:"task-frequency" desc:"default task frequency" default:"1m"`
}

type Cors struct {
	AllowOrigins []string `flag:"allow-origin" desc:"allowed origins, if not provided cors is not enabled"`
}

func (c *Config) Bind(cmd *cobra.Command, flg *pflag.FlagSet, vip *viper.Viper, name string, prefix string, keyPrefix string) {
	_ = "STUB: not implemented"
	return
}

func (c *Config) Decode(value any, decodeHook mapstructure.DecodeHookFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) New(a i_api.API, metrics *metrics.Metrics) (i_api.Subsystem, error) {
	_ = "STUB: not implemented"
	return *new(i_api.Subsystem), nil
}

type Http struct {
	config   *Config
	listen   net.Listener
	server   *http.Server
	shutdown chan struct{}
}

func New(a i_api.API, metrics *metrics.Metrics, config *Config) (i_api.Subsystem, error) {
	_ = "STUB: not implemented"
	return *new(i_api.Subsystem), nil
}

// create a shutdown channel

// Create a listener on specified address

// Register custom validators

// Middleware for logging and metrics

// CORS

// Ping endpoint, no auth required

// Authentication

// Resonate header middleware

// Extract and normalize Authorization header

// Skip "Bearer "

// Promises API

// Callbacks API
// Deprecated, use /promises/callback instead

// Subscriptions API
// Deprecated, use /promises/subscribe instead

// Schedules API

// Tasks API

// Poller proxy

// Run this request through the api middleware

func (h *Http) String() string { _ = "STUB: not implemented"; return "" }

func (h *Http) Kind() string { _ = "STUB: not implemented"; return "" }

func (h *Http) Addr() string { _ = "STUB: not implemented"; return "" }

func (h *Http) Start(errors chan<- error) {
	_ = "STUB: not implemented"
	// Start the http server
	return
}

func (h *Http) Stop() error { _ = "STUB: not implemented"; return nil }

// close the shutdown channel to immediately close all long polling
// connections all other connections remain open until complete or
// the timeout occurs

type server struct {
	api    *api.API
	config *Config
}

func (s *server) code(status t_api.StatusCode) int { _ = "STUB: not implemented"; return 0 }

// Helper functions

func oneOfCaseInsensitive(f validator.FieldLevel) bool { _ = "STUB: not implemented"; return false }
