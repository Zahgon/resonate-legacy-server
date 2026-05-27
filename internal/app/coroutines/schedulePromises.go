package coroutines

import (
	"github.com/resonatehq/gocoro"
	"github.com/resonatehq/resonate/internal/kernel/t_aio"
)

func SchedulePromises(c gocoro.Coroutine[*t_aio.Submission, *t_aio.Completion, any], m map[string]string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// read elapsed schedules

// add invocation tags

// create promise command

// Helper functions

func generatePromiseId(id string, vars map[string]string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
