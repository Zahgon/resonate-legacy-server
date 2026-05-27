package echo

import (
	"github.com/resonatehq/resonate/internal/aio"
	"github.com/resonatehq/resonate/internal/kernel/bus"
	"github.com/resonatehq/resonate/internal/kernel/t_aio"
	"github.com/resonatehq/resonate/internal/metrics"
)

// Config

type Config struct {
	Size      int `flag:"size" desc:"submission buffered channel size" default:"100"`
	BatchSize int `flag:"batch-size" desc:"max submissions processed per iteration" default:"100"`
	Workers   int `flag:"workers" desc:"number of workers" default:"1"`
}

// Subsystem

type Echo struct {
	config  *Config
	sq      chan<- *bus.SQE[t_aio.Submission, t_aio.Completion]
	workers []*EchoWorker
}

func New(aio aio.AIO, metrics *metrics.Metrics, config *Config) (*Echo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Echo) String() string { _ = "STUB: not implemented"; return "" }

func (e *Echo) Kind() t_aio.Kind { _ = "STUB: not implemented"; return *new(t_aio.Kind) }

func (e *Echo) Start(chan<- error) error { _ = "STUB: not implemented"; return nil }

func (e *Echo) Stop() error { _ = "STUB: not implemented"; return nil }

func (e *Echo) Enqueue(sqe *bus.SQE[t_aio.Submission, t_aio.Completion]) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *Echo) Flush(t int64) { _ = "STUB: not implemented"; return }

func (e *Echo) Process(sqes []*bus.SQE[t_aio.Submission, t_aio.Completion]) []*bus.CQE[t_aio.Submission, t_aio.Completion] {
	_ = "STUB: not implemented"
	return nil
}

// Worker

type EchoWorker struct {
	i       int
	sq      <-chan *bus.SQE[t_aio.Submission, t_aio.Completion]
	flush   chan int64
	aio     aio.AIO
	metrics *metrics.Metrics
}

func (w *EchoWorker) String() string { _ = "STUB: not implemented"; return "" }

func (w *EchoWorker) Start() { _ = "STUB: not implemented"; return }

// process one at a time

func (w *EchoWorker) Process(sqe *bus.SQE[t_aio.Submission, t_aio.Completion]) *bus.CQE[t_aio.Submission, t_aio.Completion] {
	_ = "STUB: not implemented"
	return nil
}

// propagate the tags
