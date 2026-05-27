package base

import (
	"time"

	"github.com/resonatehq/resonate/internal/metrics"
	"github.com/resonatehq/resonate/internal/plugins"
)

type BaseConfig struct {
	Size        int           `flag:"size" desc:"submission buffered channel size" default:"1000"`
	Workers     int           `flag:"workers" desc:"number of workers" default:"4"`
	TimeToRetry time.Duration `flag:"ttr" desc:"time to wait before resending" default:"15s"`
	TimeToClaim time.Duration `flag:"ttc" desc:"time to wait for claim before resending" default:"1m"`
}

type Processor interface {
	Process(addr []byte, head map[string]string, body []byte) (bool, error)
}

type Worker struct {
	id        int
	sq        <-chan *plugins.Message
	processor Processor
	config    *BaseConfig
	metrics   *metrics.Metrics
	name      string
}

type Plugin struct {
	name    string
	sq      chan *plugins.Message
	workers []*Worker
	cleanup func() error
}

func NewPlugin(name string, config *BaseConfig, metrics *metrics.Metrics, processor Processor, cleanup func() error) *Plugin {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) String() string { _ = "STUB: not implemented"; return "" }

func (p *Plugin) Type() string { _ = "STUB: not implemented"; return "" }

func (p *Plugin) Start(chan<- error) error { _ = "STUB: not implemented"; return nil }

func (p *Plugin) Stop() error { _ = "STUB: not implemented"; return nil }

func (p *Plugin) Enqueue(msg *plugins.Message) bool { _ = "STUB: not implemented"; return false }

func (p *Plugin) Addr() string { _ = "STUB: not implemented"; return "" }

func (w *Worker) String() string { _ = "STUB: not implemented"; return "" }

func (w *Worker) Start() { _ = "STUB: not implemented"; return }
