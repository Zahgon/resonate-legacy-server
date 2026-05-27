package coroutines

import (
	"github.com/resonatehq/gocoro"
	"github.com/resonatehq/resonate/internal/kernel/t_aio"
	"github.com/resonatehq/resonate/internal/kernel/t_api"
)

func DropTask(c gocoro.Coroutine[*t_aio.Submission, *t_aio.Completion, any], r *t_api.Request) (*t_api.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// When updating the task we don't want to reset the counter
// instead increase the counter by one in case the client tries
// to claim the same old task again. Resetting the counter
// to 1 would mean that nodes with stale tasks could claim them.

// set status

// update task

// count tasks

// It's possible that the task was modified by another coroutine
// while we were trying to complete. In that case, we should just retry.
