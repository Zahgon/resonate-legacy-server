package system

import (
	"time"

	"github.com/resonatehq/gocoro"
	"github.com/resonatehq/gocoro/pkg/promise"
	"github.com/resonatehq/resonate/internal/aio"
	"github.com/resonatehq/resonate/internal/api"
	"github.com/resonatehq/resonate/internal/metrics"

	"github.com/resonatehq/resonate/internal/kernel/t_aio"
	"github.com/resonatehq/resonate/internal/kernel/t_api"
)

type Config struct {
	Url                   string        `flag:"url" desc:"resonate server url"`
	CoroutineMaxSize      int           `flag:"coroutine-max-size" desc:"max concurrent coroutines" default:"1000" dst:"1:1000"`
	SubmissionBatchSize   int           `flag:"submission-batch-size" desc:"max submissions processed per tick" default:"1000" dst:"1:1000"`
	CompletionBatchSize   int           `flag:"completion-batch-size" desc:"max completions processed per tick" default:"1000" dst:"1:1000"`
	PromiseBatchSize      int           `flag:"promise-batch-size" desc:"max promises processed per iteration" default:"100" dst:"1:100"`
	PromiseMaxIterations  int           `flag:"promise-max-iterations" desc:"max promise iterations per coroutine" default:"1000" dst:"1:1000"`
	ScheduleBatchSize     int           `flag:"schedule-batch-size" desc:"max schedules processed per iteration" default:"100" dst:"1:100"`
	ScheduleMaxIterations int           `flag:"schedule-max-iterations" desc:"max schedule iterations per coroutine" default:"1000" dst:"1:1000"`
	TaskBatchSize         int           `flag:"task-batch-size" desc:"max tasks processed per iteration" default:"100" dst:"1:100"`
	TaskMaxIterations     int           `flag:"task-max-iterations" desc:"max task iterations per coroutine" default:"1000" dst:"1:1000"`
	SignalTimeout         time.Duration `flag:"signal-timeout" desc:"time to wait for api/aio signal" default:"1s" dst:"1s:10s"`
}

func (c *Config) String() string { _ = "STUB: not implemented"; return "" }

type backgroundCoroutine struct {
	coroutine func(map[string]string) gocoro.CoroutineFunc[*t_aio.Submission, *t_aio.Completion, any]
	name      string
	last      int64
	promise   promise.Promise[any]
}

type System struct {
	api          api.API
	aio          aio.AIO
	config       *Config
	metrics      *metrics.Metrics
	scheduler    gocoro.Scheduler[*t_aio.Submission, *t_aio.Completion]
	onRequest    map[t_api.Kind]func(req *t_api.Request, res func(*t_api.Response, error)) gocoro.CoroutineFunc[*t_aio.Submission, *t_aio.Completion, any]
	background   []*backgroundCoroutine
	shutdown     chan any
	shortCircuit chan any
}

func New(api api.API, aio aio.AIO, config *Config, metrics *metrics.Metrics) *System {
	_ = "STUB: not implemented"
	return nil
}

func (s *System) String() string { _ = "STUB: not implemented"; return "" }

func (s *System) Loop() error { _ = "STUB: not implemented"; return nil }

// tick first

// complete shutdown if done

// create signals

// wait for a signal, short circuit, or timeout; whichever occurs
// first

// close the cancel channel so the api and aio can stop listening
// and wait for the signal channels to close

func (s *System) Tick(t int64) { _ = "STUB: not implemented"; return }

// dequeue sqes and cqes

// call completion callbacks

// add background coroutines

// system is shutting down

// background coroutines are mutually exclusive

// wait min amount of time between scheduling

// add request coroutines

// tick scheduler

// flush aio

func (s *System) Shutdown() <-chan any {
	_ = "STUB: not implemented"
	// start by shutting down the api
	return nil
}

// short circuit the system loop

// return the shutdown channel so the caller can wait on system
// shutdown

func (s *System) Done() bool { _ = "STUB: not implemented"; return false }

func (s *System) AddOnRequest(kind t_api.Kind, constructor func(gocoro.Coroutine[*t_aio.Submission, *t_aio.Completion, any], *t_api.Request) (*t_api.Response, error)) {
	_ = "STUB: not implemented"
	return
}

// set config

// set metrics

// run coroutine

func (s *System) AddBackground(name string, constructor func(gocoro.Coroutine[*t_aio.Submission, *t_aio.Completion, any], map[string]string) (any, error)) {
	_ = "STUB: not implemented"
	return
}

// set config

// set metrics

// run coroutine

// TODO: move this to gocoro
func (s *System) coroutineMetrics(p promise.Promise[any], tags map[string]string) {
	_ = "STUB: not implemented"
	return
}
