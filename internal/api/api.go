package api

import (
	"github.com/resonatehq/resonate/internal/kernel/bus"
	"github.com/resonatehq/resonate/internal/kernel/t_api"
	"github.com/resonatehq/resonate/internal/metrics"
	"github.com/resonatehq/resonate/internal/plugins"
)

type API interface {
	String() string

	Start() error
	Stop() error
	Shutdown()
	Done() bool
	Errors() <-chan error

	Signal(<-chan interface{}) <-chan interface{}

	Plugins() []plugins.Plugin

	EnqueueSQE(*bus.SQE[t_api.Request, t_api.Response])
	DequeueSQE(int) []*bus.SQE[t_api.Request, t_api.Response]
	EnqueueCQE(*bus.CQE[t_api.Request, t_api.Response])
	DequeueCQE(<-chan *bus.CQE[t_api.Request, t_api.Response]) *bus.CQE[t_api.Request, t_api.Response]
}

// API

type api struct {
	sq         chan *bus.SQE[t_api.Request, t_api.Response]
	buffer     *bus.SQE[t_api.Request, t_api.Response]
	subsystems []Subsystem
	plugins    []plugins.Plugin
	done       bool
	errors     chan error
	metrics    *metrics.Metrics
	middleware []Middleware
}

func New(size int, metrics *metrics.Metrics) *api { _ = "STUB: not implemented"; return nil }

func (a *api) String() string { _ = "STUB: not implemented"; return "" }

func (a *api) AddSubsystem(subsystem Subsystem) { _ = "STUB: not implemented"; return }

func (a *api) AddPlugin(plugin plugins.Plugin) { _ = "STUB: not implemented"; return }

func (a *api) Plugins() []plugins.Plugin { _ = "STUB: not implemented"; return nil }

func (a *api) AddMiddleware(middleware Middleware) { _ = "STUB: not implemented"; return }

func (a *api) Addr() string {
	_ = "STUB: not implemented"
	// advertise only the http address, if available
	return ""
}

// Lifecycle functions

func (a *api) Start() error { _ = "STUB: not implemented"; return nil }

func (a *api) Stop() error { _ = "STUB: not implemented"; return nil }

func (a *api) Shutdown() { _ = "STUB: not implemented"; return }

func (a *api) Done() bool { _ = "STUB: not implemented"; return false }

func (a *api) Errors() <-chan error {
	_ = "STUB: not implemented"

	// IO functions
	return nil
}

func (a *api) Signal(cancel <-chan any) <-chan any { _ = "STUB: not implemented"; return nil }

// SQE

func (a *api) EnqueueSQE(sqe *bus.SQE[t_api.Request, t_api.Response]) {
	_ = "STUB: not implemented"
	return
}

// replace callback with a function that emits metrics

// we must wait to close the channel because even in a select
// sending to a closed channel will panic

// validate the submission before sending it to the cq, this will
// run in the same goroutine as the api request in order to fail
// fast

// Run all the middleware

func (a *api) DequeueSQE(n int) []*bus.SQE[t_api.Request, t_api.Response] {
	_ = "STUB: not implemented"
	return nil
}

// insert the buffered sqe

// collects n entries (if immediately available)

// CQE

func (a *api) EnqueueCQE(cqe *bus.CQE[t_api.Request, t_api.Response]) {
	_ = "STUB: not implemented"
	return
}

func (a *api) DequeueCQE(cq <-chan *bus.CQE[t_api.Request, t_api.Response]) *bus.CQE[t_api.Request, t_api.Response] {
	_ = "STUB: not implemented"
	return nil
}
