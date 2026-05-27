package dst

import (
	"math/rand" // nosemgrep
	"time"

	"github.com/anishathalye/porcupine"
	"github.com/resonatehq/resonate/internal/aio"
	"github.com/resonatehq/resonate/internal/api"
	"github.com/resonatehq/resonate/internal/kernel/system"
	"github.com/resonatehq/resonate/internal/kernel/t_api"
	"github.com/resonatehq/resonate/pkg/promise"
	"github.com/resonatehq/resonate/pkg/task"
)

type DST struct {
	config      *Config
	generator   *Generator
	validator   *Validator
	bcValidator *BcValidator
	partitions  [][]porcupine.Operation // set by Partition in the porcupine model
}

type Config struct {
	Ticks              int64
	Timeout            time.Duration
	VisualizationPath  string
	Verbose            bool
	PrintOps           bool
	TimeElapsedPerTick int64
	TimeoutTicks       int64
	ReqsPerTick        func() int
	MaxReqsPerTick     int64
	Ids                int
	Headers            int
	Data               int
	Tags               int
	FaultInjection     bool
	Backchannel        chan interface{}
}

type Kind int

const (
	Op Kind = iota
	Bc
)

type Partition int

type Req struct {
	kind Kind
	time int64
	req  *t_api.Request
	bc   *Backchannel
}

type Res struct {
	kind Kind
	time int64
	res  *t_api.Response
	err  error
}

type BcKind int

const (
	Task BcKind = iota
	Notify
)

type Backchannel struct {
	Task    *task.Task
	Promise *promise.Promise
}

func New(r *rand.Rand, config *Config) *DST { _ = "STUB: not implemented"; return nil }

func (d *DST) Add(kind t_api.Kind, generator RequestGenerator, validator ResponseValidator) {
	_ = "STUB: not implemented"
	return
}

func (d *DST) Run(r *rand.Rand, api api.API, aio aio.AIO, system *system.System) bool {
	_ = "STUB: not implemented"
	return false
}

// promises

// callbacks

// schedules

// tasks

// backchannel validators

// porcupine ops

// run all requests through the server and collect responses

// subtract 1 to ensure tick timeframes don't overlap

// log

// add operation to porcupine

// now read from the backchannel

// skip scheduled promises, this is a little hacky but we know
// that scheduled promises start with an 's'

// randomly decrement the counter, we only decrement so that we
// know a successful claim task request can only occur after
// our model has been updated via the backchannel

// The processId is always the taskId, which means each task
// is always claimed by a unique process. When heartbeating there
// will be most a single task.

// add claim req to generator

// skip scheduled promises, this is a little hacky but we know
// that scheduled promises start with an 's'

// backchannel messages occur on the "last" tick

// add backchannel op to porcupine

// shutdown the system

// keep ticking until all submissions have been processed

func (d *DST) logPossibleError(history porcupine.LinearizationInfo) {
	_ = "STUB: not implemented"
	// Whats is printed here and whats is visualized in the dst.html diagram might not match.
	// this is a best effort to preserve the possible validation that failed.
	return
}

// check each parition individually
// partitions are in the order they were given to porcupine

// take the first (and we assume, by empiric evidence, only linearization)

// if the linearization includes all the operations in the partiton all good

func (d *DST) logError(partialLinearization []porcupine.Operation, lastOp porcupine.Operation) {
	_ = "STUB: not implemented"
	// create a new model
	return
}

// re feed operations through model

// step through the model (again)

func (d *DST) Model() porcupine.Model { _ = "STUB: not implemented"; return *new(porcupine.Model) }

// Get sorted keys to iterate over the partitions in a deterministic way

func (d *DST) Step(model *Model, reqTime int64, resTime int64, req *t_api.Request, res *t_api.Response, err error) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sometimes we generate create callback requests with the same root and
// leaf promise ids by chance

func (d *DST) StepBc(model *Model, reqTime int64, resTime int64, req *Req) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DST) Time(t int64) int64 { _ = "STUB: not implemented"; return 0 }

func (d *DST) String() string { _ = "STUB: not implemented"; return "" }

// Helper functions

func partition(req *Req) string { _ = "STUB: not implemented"; return "" }

// Find the first Operation if any that is not part of a partial linearization
// by comparing our partition with the linearization
func nextFailure(linearizationOps []porcupine.Operation, partitionOps []porcupine.Operation) porcupine.Operation {
	_ = "STUB: not implemented"
	// convert to map for quick lookup
	return *new(porcupine.Operation)
}

// if req is part of the linearizable path, skip

// ops are ordered by time, so the first op is not part of the
// linearizable path should break the model
