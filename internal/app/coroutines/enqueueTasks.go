package coroutines

import (
	"github.com/resonatehq/gocoro"
	"github.com/resonatehq/resonate/internal/kernel/t_aio"
)

func EnqueueTasks(c gocoro.Coroutine[*t_aio.Submission, *t_aio.Completion, any], m map[string]string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// go straight to jail, do not collect $200

// also go straight to jail, do not collect $200

// time to claim

// fallback to 15s

// time to retry enqueueing
