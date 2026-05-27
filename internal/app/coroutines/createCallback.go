package coroutines

import (
	"github.com/resonatehq/gocoro"
	"github.com/resonatehq/resonate/internal/kernel/t_aio"
	"github.com/resonatehq/resonate/internal/kernel/t_api"
)

func CreateCallback(c gocoro.Coroutine[*t_aio.Submission, *t_aio.Completion, any], r *t_api.Request) (*t_api.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// read the promise to see if it exists

// If the callback was already created return 200 and an empty callback

// Status could be StatusOk or StatusCreated if the Callback Id was already present
