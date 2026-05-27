package test

import (
	"testing"

	"github.com/resonatehq/resonate/internal/kernel/bus"
	"github.com/resonatehq/resonate/internal/kernel/t_api"
	"github.com/resonatehq/resonate/internal/plugins"
)

type API struct {
	t   *testing.T
	req *t_api.Request
	res *t_api.Response
}

func (a *API) Load(t *testing.T, req *t_api.Request, res *t_api.Response) {
	_ = "STUB: not implemented"
	return
}

func (a *API) String() string { _ = "STUB: not implemented"; return "" }

func (a *API) Start() error { _ = "STUB: not implemented"; return nil }

func (a *API) Stop() error { _ = "STUB: not implemented"; return nil }

func (a *API) Shutdown() { _ = "STUB: not implemented"; return }

func (a *API) Done() bool { _ = "STUB: not implemented"; return false }

func (a *API) Errors() <-chan error { _ = "STUB: not implemented"; return nil }

func (a *API) Signal(cancel <-chan interface{}) <-chan interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) Plugins() []plugins.Plugin { _ = "STUB: not implemented"; return nil }

func (a *API) EnqueueSQE(sqe *bus.SQE[t_api.Request, t_api.Response]) {
	_ = "STUB: not implemented"
	// assert
	return
}

// immediately call callback

func (a *API) DequeueSQE(int) []*bus.SQE[t_api.Request, t_api.Response] {
	_ = "STUB: not implemented"
	return nil
}

func (a *API) EnqueueCQE(*bus.CQE[t_api.Request, t_api.Response]) {
	_ = "STUB: not implemented"
	return
}

func (a *API) DequeueCQE(cq <-chan *bus.CQE[t_api.Request, t_api.Response]) *bus.CQE[t_api.Request, t_api.Response] {
	_ = "STUB: not implemented"
	return nil
}
