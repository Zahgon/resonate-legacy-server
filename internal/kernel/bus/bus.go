package bus

import (
	"github.com/resonatehq/resonate/internal/kernel/t_aio"
	"github.com/resonatehq/resonate/internal/kernel/t_api"
)

type Input interface {
	t_aio.Submission | t_api.Request
}

type Output interface {
	t_aio.Completion | t_api.Response
}

type SQE[I Input, O Output] struct {
	Id         string
	Callback   func(*O, error)
	Submission *I
}

func (sqe *SQE[I, O]) String() string { _ = "STUB: not implemented"; return "" }

type CQE[I Input, O Output] struct {
	Id         string
	Callback   func(*O, error)
	Completion *O
	Error      error
}

func (cqe *CQE[I, O]) Invoke() { _ = "STUB: not implemented"; return }

func (cqe *CQE[I, O]) String() string { _ = "STUB: not implemented"; return "" }
