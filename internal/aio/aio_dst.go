package aio

import (
	"math/rand" // nosemgrep

	"github.com/resonatehq/resonate/internal/kernel/t_aio"
	"github.com/resonatehq/resonate/internal/metrics"
	"github.com/resonatehq/resonate/internal/plugins"

	"github.com/resonatehq/resonate/internal/kernel/bus"
)

type aioDST struct {
	r          *rand.Rand
	p          float64
	sqes       []*bus.SQE[t_aio.Submission, t_aio.Completion]
	cqes       []*bus.CQE[t_aio.Submission, t_aio.Completion]
	subsystems map[t_aio.Kind]SubsystemDST
	metrics    *metrics.Metrics
}

func NewDST(r *rand.Rand, p float64, metrics *metrics.Metrics) *aioDST {
	_ = "STUB: not implemented"
	return nil
}

func (a *aioDST) String() string {
	_ = "STUB: not implemented"
	// use subsystem keys so that we can compare cross-store dst runs
	return ""
}

func (a *aioDST) AddSubsystem(subsystem SubsystemDST) { _ = "STUB: not implemented"; return }

func (a *aioDST) Plugins() []plugins.Plugin { _ = "STUB: not implemented"; return nil }

func (a *aioDST) Start() error { _ = "STUB: not implemented"; return nil }

func (a *aioDST) Stop() error { _ = "STUB: not implemented"; return nil }

func (a *aioDST) Shutdown() { _ = "STUB: not implemented"; return }

func (a *aioDST) Errors() <-chan error { _ = "STUB: not implemented"; return nil }

func (a *aioDST) Signal(cancel <-chan interface{}) <-chan interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (a *aioDST) Flush(t int64) { _ = "STUB: not implemented"; return }

// there is a p percent chance of failure
// either pre or post processing

// simulate failure before processing

// process the SQE

// simulate failure after processing

// SQE

func (a *aioDST) Dispatch(submission *t_aio.Submission, callback func(*t_aio.Completion, error)) {
	_ = "STUB: not implemented"
	return
}

func (a *aioDST) EnqueueSQE(sqe *bus.SQE[t_aio.Submission, t_aio.Completion]) {
	_ = "STUB: not implemented"
	return
}

// insert at random position

// CQE

func (a *aioDST) EnqueueCQE(cqe *bus.CQE[t_aio.Submission, t_aio.Completion]) {
	_ = "STUB: not implemented"
	return
}

func (a *aioDST) DequeueCQE(n int) []*bus.CQE[t_aio.Submission, t_aio.Completion] {
	_ = "STUB: not implemented"
	return nil
}
