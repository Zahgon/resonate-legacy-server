package sender

import (
	"encoding/json"
	"math/rand" // nosemgrep

	"github.com/go-viper/mapstructure/v2"
	"github.com/resonatehq/resonate/internal/aio"
	"github.com/resonatehq/resonate/internal/kernel/bus"
	"github.com/resonatehq/resonate/internal/kernel/t_aio"
	"github.com/resonatehq/resonate/internal/metrics"
	"github.com/resonatehq/resonate/internal/plugins"
	"github.com/resonatehq/resonate/pkg/receiver"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config

type Config struct {
	Size    int            `flag:"size" desc:"submission buffered channel size" default:"100"`
	Targets []TargetConfig `flag:"targets" desc:"target config"`
}

type TargetConfig struct {
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

func (c *Config) NewDST(aio.AIO, *metrics.Metrics, *rand.Rand, chan interface{}) (aio.SubsystemDST, error) {
	_ = "STUB: not implemented"
	return *

	// Subsystem
	new(aio.SubsystemDST), nil
}

type Sender struct {
	config *Config
	sq     chan<- *bus.SQE[t_aio.Submission, t_aio.Completion]
	worker *SenderWorker
}

func New(a aio.AIO, metrics *metrics.Metrics, config *Config) (*Sender, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// add default target if none

func (s *Sender) String() string { _ = "STUB: not implemented"; return "" }

func (s *Sender) Kind() t_aio.Kind { _ = "STUB: not implemented"; return *new(t_aio.Kind) }

func (s *Sender) Start(errors chan<- error) error {
	_ = "STUB: not implemented"
	// start worker
	return nil
}

func (s *Sender) Stop() error { _ = "STUB: not implemented"; return nil }

func (s *Sender) Enqueue(sqe *bus.SQE[t_aio.Submission, t_aio.Completion]) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Sender) Flush(t int64) {
	_ = "STUB: not implemented"

	// Worker
	return
}

type SenderWorker struct {
	sq      <-chan *bus.SQE[t_aio.Submission, t_aio.Completion]
	plugins map[string]plugins.Plugin
	targets map[string]*receiver.Recv
	aio     aio.AIO
	metrics *metrics.Metrics
}

func (w *SenderWorker) String() string { _ = "STUB: not implemented"; return "" }

func (w *SenderWorker) AddPlugin(plugin plugins.Plugin) { _ = "STUB: not implemented"; return }

func (w *SenderWorker) Start() { _ = "STUB: not implemented"; return }

// process one at a time

func (w *SenderWorker) Process(sqe *bus.SQE[t_aio.Submission, t_aio.Completion]) {
	_ = "STUB: not implemented"
	return
}

// instantiate cqe

func boolToStatus(b bool) string { _ = "STUB: not implemented"; return "" }

func schemeToRecv(v string) (*receiver.Recv, bool) { _ = "STUB: not implemented"; return nil, false }
