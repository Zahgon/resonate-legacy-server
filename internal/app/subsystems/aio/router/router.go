package router

import (
	"encoding/json"
	"math/rand" // nosemgrep

	"github.com/go-viper/mapstructure/v2"
	"github.com/resonatehq/resonate/internal/aio"
	"github.com/resonatehq/resonate/internal/kernel/bus"
	"github.com/resonatehq/resonate/internal/kernel/t_aio"
	"github.com/resonatehq/resonate/internal/metrics"
	"github.com/resonatehq/resonate/pkg/promise"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Size    int            `flag:"size" desc:"submission buffered channel size" default:"100"`
	Workers int            `flag:"workers" desc:"number of workers" default:"1" dst:"1"`
	Sources []SourceConfig `flag:"sources" desc:"source config"`
}

type SourceConfig struct {
	Name string
	Type string
	Data json.RawMessage
}

func (c *Config) Bind(cmd *cobra.Command, flg *pflag.FlagSet, vip *viper.Viper, name string, prefix string, keyPrefix string) {
	_ = "STUB: not implemented"
	return
}

func (c *Config) Decode(value any, decodeHook mapstructure.DecodeHookFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) New(aio aio.AIO, metrics *metrics.Metrics) (aio.Subsystem, error) {
	_ = "STUB: not implemented"
	return *new(aio.Subsystem), nil
}

func (c *Config) NewDST(aio aio.AIO, metrics *metrics.Metrics, _ *rand.Rand, _ chan any) (aio.SubsystemDST, error) {
	_ = "STUB: not implemented"
	return *new(aio.SubsystemDST), nil
}

type TagSourceConfig struct {
	Key string
}

// Subsystem

type Router struct {
	config  *Config
	sq      chan<- *bus.SQE[t_aio.Submission, t_aio.Completion]
	workers []*RouterWorker
}

func New(aio aio.AIO, metrics *metrics.Metrics, config *Config) (*Router, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// add default source if none found

func (r *Router) String() string { _ = "STUB: not implemented"; return "" }

func (r *Router) Kind() t_aio.Kind { _ = "STUB: not implemented"; return *new(t_aio.Kind) }

func (r *Router) Start(chan<- error) error { _ = "STUB: not implemented"; return nil }

func (r *Router) Stop() error { _ = "STUB: not implemented"; return nil }

func (r *Router) Enqueue(sqe *bus.SQE[t_aio.Submission, t_aio.Completion]) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Router) Flush(int64) { _ = "STUB: not implemented"; return }

func (r *Router) Process(sqes []*bus.SQE[t_aio.Submission, t_aio.Completion]) []*bus.CQE[t_aio.Submission, t_aio.Completion] {
	_ = "STUB: not implemented"
	return nil
}

// Worker

type RouterWorker struct {
	i       int
	sq      <-chan *bus.SQE[t_aio.Submission, t_aio.Completion]
	sources []func(*promise.Promise) (any, bool)
	aio     aio.AIO
	metrics *metrics.Metrics
}

func (w *RouterWorker) String() string { _ = "STUB: not implemented"; return "" }

func (w *RouterWorker) Start() { _ = "STUB: not implemented"; return }

// process one at a time

func (w *RouterWorker) Process(sqe *bus.SQE[t_aio.Submission, t_aio.Completion]) *bus.CQE[t_aio.Submission, t_aio.Completion] {
	_ = "STUB: not implemented"
	return nil
}

// apply sources in succession, first match wins

// Source functions

func TagSource(config *TagSourceConfig) func(*promise.Promise) (any, bool) {
	_ = "STUB: not implemented"
	return nil
}

// not a match if tag is not present

// check if tag is valid json

// valid json is a match

// invalid json is not a match

// otherwise return the string value

// Helper functions

func coerce(v any) (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

// will be matched against logical receivers in the queueing
// subsystem
