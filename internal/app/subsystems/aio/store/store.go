package store

import (
	"github.com/resonatehq/resonate/internal/kernel/bus"
	"github.com/resonatehq/resonate/internal/kernel/t_aio"
)

type Store interface {
	Execute([]*t_aio.Transaction) ([]*t_aio.StoreCompletion, error)
}

func Process(store Store, sqes []*bus.SQE[t_aio.Submission, t_aio.Completion]) []*bus.CQE[t_aio.Submission, t_aio.Completion] {
	_ = "STUB: not implemented"
	return nil
}

// propagate the tags

func Collect(c <-chan *bus.SQE[t_aio.Submission, t_aio.Completion], f <-chan int64, n int) ([]*bus.SQE[t_aio.Submission, t_aio.Completion], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func StoreErr(err error) error { _ = "STUB: not implemented"; return nil }
