package aio

import (
	"github.com/resonatehq/resonate/internal/kernel/bus"
	"github.com/resonatehq/resonate/internal/kernel/t_aio"
	"github.com/resonatehq/resonate/internal/plugins"

	"github.com/resonatehq/resonate/internal/metrics"
)

type AIO interface {
	String() string

	Start() error
	Stop() error
	Shutdown()
	Errors() <-chan error

	Signal(<-chan interface{}) <-chan interface{}
	Flush(int64)

	Plugins() []plugins.Plugin

	// dispatch is required by gocoro
	Dispatch(*t_aio.Submission, func(*t_aio.Completion, error))

	EnqueueSQE(*bus.SQE[t_aio.Submission, t_aio.Completion])
	EnqueueCQE(*bus.CQE[t_aio.Submission, t_aio.Completion])
	DequeueCQE(int) []*bus.CQE[t_aio.Submission, t_aio.Completion]
}

// AIO

type aio struct {
	cq         chan *bus.CQE[t_aio.Submission, t_aio.Completion]
	buffer     *bus.CQE[t_aio.Submission, t_aio.Completion]
	subsystems map[t_aio.Kind]Subsystem
	plugins    []plugins.Plugin
	errors     chan error
	metrics    *metrics.Metrics
}

func New(size int, metrics *metrics.Metrics) *aio { _ = "STUB: not implemented"; return nil }

func (a *aio) String() string { _ = "STUB: not implemented"; return "" }

func (a *aio) AddSubsystem(subsystem Subsystem) { _ = "STUB: not implemented"; return }

func (a *aio) AddPlugin(plugin plugins.Plugin) { _ = "STUB: not implemented"; return }

func (a *aio) Plugins() []plugins.Plugin {
	_ = "STUB: not implemented"

	// Lifecycle functions
	return nil
}

func (a *aio) Start() error {
	_ = "STUB: not implemented"
	// start plugins
	return nil
}

// start subsystems

func (a *aio) Stop() error {
	_ = "STUB: not implemented"

	// stop subsystems
	return nil
}

// stop plugins

func (a *aio) Shutdown() { _ = "STUB: not implemented"; return }

func (a *aio) Errors() <-chan error {
	_ = "STUB: not implemented"

	// IO functions
	return nil
}

func (a *aio) Signal(cancel <-chan any) <-chan any { _ = "STUB: not implemented"; return nil }

func (a *aio) Flush(t int64) { _ = "STUB: not implemented"; return }

// SQE

func (a *aio) Dispatch(submission *t_aio.Submission, callback func(*t_aio.Completion, error)) {
	_ = "STUB: not implemented"
	return
}

func (a *aio) EnqueueSQE(sqe *bus.SQE[t_aio.Submission, t_aio.Completion]) {
	_ = "STUB: not implemented"
	return
}

// CQE

func (a *aio) EnqueueCQE(cqe *bus.CQE[t_aio.Submission, t_aio.Completion]) {
	_ = "STUB: not implemented"
	return
}

// block until the completion queue has space

func (a *aio) DequeueCQE(n int) []*bus.CQE[t_aio.Submission, t_aio.Completion] {
	_ = "STUB: not implemented"
	return nil
}

// insert the buffered sqe

// collects n entries (if immediately available)
