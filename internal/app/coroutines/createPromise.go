package coroutines

import (
	"github.com/resonatehq/gocoro"
	"github.com/resonatehq/resonate/internal/kernel/t_aio"
	"github.com/resonatehq/resonate/internal/kernel/t_api"
	"github.com/resonatehq/resonate/pkg/promise"
	"github.com/resonatehq/resonate/pkg/task"
)

func CreatePromise(c gocoro.Coroutine[*t_aio.Submission, *t_aio.Completion, any], r *t_api.Request) (*t_api.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// determine initial state

func CreatePromiseAndTask(c gocoro.Coroutine[*t_aio.Submission, *t_aio.Completion, any], r *t_api.Request) (*t_api.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// determine initial state

type promiseAndTask struct {
	created bool
	promise *promise.Promise
	task    *task.Task
}

func createPromise(tags map[string]string, promiseCmd *t_aio.CreatePromiseCommand, taskCmd *t_aio.CreateTaskCommand, additionalCmds ...t_aio.Command) gocoro.CoroutineFunc[*t_aio.Submission, *t_aio.Completion, *promiseAndTask] {
	_ = "STUB: not implemented"
	return nil
}

// first read the promise to see if it already exists

// response data

// first, check the router to see if a task needs to be created

// just set the recv

// add the task command

// add a create promise and task command

// add create promise command

// It's possible that the promise was created by another coroutine
// while we were creating. In that case, we should just retry.

// count promise

// count task (if applicable)

// It's possible that the promise was created by another coroutine
// while we were creating. In that case, we should just retry.

// update promise
