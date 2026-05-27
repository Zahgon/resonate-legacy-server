package api

import (
	i_api "github.com/resonatehq/resonate/internal/api"
	"github.com/resonatehq/resonate/internal/kernel/t_api"
)

type API struct {
	i_api.API
	protocol string
}

func New(api i_api.API, protocol string) *API { _ = "STUB: not implemented"; return nil }

func (a *API) Process(id string, submission *t_api.Request) (*t_api.Response, *Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// inject tags

// completion queue

// cqe := <-cq

// Helper functions

func (a *API) SearchPromises(id string, state string, tags map[string]string, limit int, cursor string) (*t_api.PromiseSearchRequest, *Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validate id

// set states

// set default tags

// set default limit

// validate limit

func (a *API) SearchSchedules(id string, tags map[string]string, limit int, cursor string) (*t_api.ScheduleSearchRequest, *Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validate id

// set default tags

// set default limit

// validate limit

func (a *API) ValidateCron(cron string) *Error { _ = "STUB: not implemented"; return nil }

func (a *API) TaskProcessId(id string, counter int) string { _ = "STUB: not implemented"; return "" }
