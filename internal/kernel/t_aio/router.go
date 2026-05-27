package t_aio

import (
	"github.com/resonatehq/resonate/pkg/promise"
)

type RouterSubmission struct {
	Promise *promise.Promise
}

func (s *RouterSubmission) String() string { _ = "STUB: not implemented"; return "" }

type RouterCompletion struct {
	Matched bool
	Recv    []byte
}

func (c *RouterCompletion) String() string { _ = "STUB: not implemented"; return "" }
