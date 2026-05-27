package http

import (
	"net/http"
	"time"

	"github.com/go-viper/mapstructure/v2"
	"github.com/resonatehq/resonate/internal/app/plugins/base"
	"github.com/resonatehq/resonate/internal/metrics"
	"github.com/resonatehq/resonate/internal/plugins"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Size        int           `flag:"size" desc:"submission buffered channel size" default:"100"`
	Workers     int           `flag:"workers" desc:"number of workers" default:"3"`
	Timeout     time.Duration `flag:"timeout" desc:"http request timeout" default:"3m"`
	ConnTimeout time.Duration `flag:"conn-timeout" desc:"http connection timeout" default:"10s"`
	TimeToRetry time.Duration `flag:"ttr" desc:"time to wait before resending" default:"15s"`
	TimeToClaim time.Duration `flag:"ttc" desc:"time to wait for claim before resending" default:"1m"`
}

func (c *Config) Bind(cmd *cobra.Command, flg *pflag.FlagSet, vip *viper.Viper, name string, prefix string, keyPrefix string) {
	_ = "STUB: not implemented"
	return
}

func (c *Config) Decode(value any, decodeHook mapstructure.DecodeHookFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) New(metrics *metrics.Metrics) (plugins.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(plugins.Plugin), nil
}

type Http struct {
	*base.Plugin
}

type Addr struct {
	Headers map[string]string `json:"headers,omitempty"`
	Url     string            `json:"url"`
}

type processor struct {
	client *http.Client
}

func New(metrics *metrics.Metrics, config *Config) (*Http, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *processor) Process(data []byte, head map[string]string, body []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// nosemgrep: range-over-map

// nosemgrep: range-over-map

// set non-overridable headers

// when a connection is established, close the channel and return
// true

// perform request asynchronously

// when request fails and the connection was not established, close
// the channel and return false
