package dst

import (
	"math/rand" // nosemgrep

	"github.com/resonatehq/resonate/internal/kernel/t_api"
)

type Generator struct {
	ticks              int64
	timeElapsedPerTick int64
	timeoutTicks       int64 // max ticks in the future to set promise timeout
	idSet              []string
	headersSet         []map[string]string
	dataSet            [][]byte
	tagsSet            []map[string]string
	requests           map[t_api.Kind][]*t_api.Request
	generators         []RequestGenerator
}

type RequestGenerator func(*rand.Rand, int64) *t_api.Request

func NewGenerator(r *rand.Rand, config *Config) *Generator { _ = "STUB: not implemented"; return nil }

// pad ids with leading zeros to ensure ids are the same length
// this helps with lexigraphical sorting across different databases

// half of all headers are nil

// half of all values are nil

// transition to resolved on timeout

// create a task

// half of all tags are nil

func (g *Generator) AddGenerator(kind t_api.Kind, generator RequestGenerator) {
	_ = "STUB: not implemented"
	return
}

func (g *Generator) AddRequest(req *t_api.Request) { _ = "STUB: not implemented"; return }

func (g *Generator) Generate(r *rand.Rand, t int64, n int) []*t_api.Request {
	_ = "STUB: not implemented"
	return nil
}

// some generators require "canned requests" and return nil if none
// are available, loop until we get a non-nil request

// PROMISES

func (g *Generator) GenerateReadPromise(r *rand.Rand, t int64) *t_api.Request {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) GenerateSearchPromises(r *rand.Rand, t int64) *t_api.Request {
	_ = "STUB: not implemented"
	// grab a cursor if one is available
	return nil
}

// states

func (g *Generator) GenerateCreatePromise(r *rand.Rand, t int64) *t_api.Request {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) GenerateCreatePromiseAndTask(r *rand.Rand, t int64) *t_api.Request {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) GenerateCompletePromise(r *rand.Rand, t int64) *t_api.Request {
	_ = "STUB: not implemented"
	return nil
}

// CALLBACKS

func (g *Generator) GenerateCreateCallback(r *rand.Rand, t int64) *t_api.Request {
	_ = "STUB: not implemented"
	return nil
}

// ignored in dst, use hardcoded value

// ignored in dst, use hardcoded value

// SCHEDULES

func (g *Generator) GenerateReadSchedule(r *rand.Rand, t int64) *t_api.Request {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) GenerateSearchSchedules(r *rand.Rand, t int64) *t_api.Request {
	_ = "STUB: not implemented"
	// grab a cursor if one is available
	return nil
}

func (g *Generator) GenerateCreateSchedule(r *rand.Rand, t int64) *t_api.Request {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) GenerateDeleteSchedule(r *rand.Rand, t int64) *t_api.Request {
	_ = "STUB: not implemented"
	return nil
}

// TASKS

func (g *Generator) GenerateClaimTask(r *rand.Rand, t int64) *t_api.Request {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) GenerateCompleteTask(r *rand.Rand, t int64) *t_api.Request {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) GenerateDropTask(r *rand.Rand, t int64) *t_api.Request {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) GenerateHeartbeatTasks(r *rand.Rand, t int64) *t_api.Request {
	_ = "STUB: not implemented"
	return nil
}

// Helpers

func (g *Generator) promiseId(r *rand.Rand) string { _ = "STUB: not implemented"; return "" }

func (g *Generator) scheduleId(r *rand.Rand) string { _ = "STUB: not implemented"; return "" }

func (g *Generator) callbackId(r *rand.Rand) string { _ = "STUB: not implemented"; return "" }

func (g *Generator) promiseSearch(r *rand.Rand) string { _ = "STUB: not implemented"; return "" }

func (g *Generator) scheduleSearch(r *rand.Rand) string { _ = "STUB: not implemented"; return "" }

func (g *Generator) search(r *rand.Rand) string { _ = "STUB: not implemented"; return "" }

func (g *Generator) pop(r *rand.Rand, kind t_api.Kind) *t_api.Request {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) nextTasks(r *rand.Rand, id string, pid string, counter int, partitionId string) {
	_ = "STUB: not implemented"
	// seed the "next" requests,
	// sometimes we deliberately do nothing
	return
}

// do nothing

func (g *Generator) tags(r *rand.Rand) map[string]string { _ = "STUB: not implemented"; return nil }

func (g *Generator) headers(r *rand.Rand) map[string]string { _ = "STUB: not implemented"; return nil }
