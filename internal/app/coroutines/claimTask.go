package coroutines

import (
	"github.com/resonatehq/gocoro"
	"github.com/resonatehq/resonate/internal/kernel/t_aio"
	"github.com/resonatehq/resonate/internal/kernel/t_api"
)

func ClaimTask(c gocoro.Coroutine[*t_aio.Submission, *t_aio.Completion, any], r *t_api.Request) (*t_api.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// guard against overflow

// time to expire unless heartbeated

// set promises

// set status

// update task

// It's possible that the task was modified by another coroutine
// while we were trying to claim. In that case, we should just retry.
