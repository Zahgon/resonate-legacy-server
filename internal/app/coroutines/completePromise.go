package coroutines

import (
	"github.com/resonatehq/gocoro"
	"github.com/resonatehq/resonate/internal/kernel/t_aio"
	"github.com/resonatehq/resonate/internal/kernel/t_api"
)

func CompletePromise(c gocoro.Coroutine[*t_aio.Submission, *t_aio.Completion, any], r *t_api.Request) (*t_api.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// It's possible that the promise was completed by another coroutine
// while we were completing. In that case, we should just retry.

func completePromise(tags map[string]string, updatePromiseCmd *t_aio.UpdatePromiseCommand, additionalCmds ...t_aio.Command) gocoro.CoroutineFunc[*t_aio.Submission, *t_aio.Completion, bool] {
	_ = "STUB: not implemented"
	return nil
}

// add additional commands

// count promises

// count tasks
