package coroutines

import (
	"github.com/resonatehq/gocoro"
	"github.com/resonatehq/resonate/internal/kernel/t_aio"
)

func TimeoutTasks(c gocoro.Coroutine[*t_aio.Submission, *t_aio.Completion, any], m map[string]string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// time until reenqueued
