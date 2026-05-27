package t_aio

import (
	"github.com/resonatehq/resonate/pkg/promise"
	"github.com/resonatehq/resonate/pkg/task"
)

type SenderSubmission struct {
	Task          *task.Task
	Promise       *promise.Promise
	BaseHref      string
	ClaimHref     string
	CompleteHref  string
	HeartbeatHref string
}

func (s *SenderSubmission) String() string { _ = "STUB: not implemented"; return "" }

type SenderCompletion struct {
	Success     bool
	TimeToRetry int64
	TimeToClaim int64
}

func (c *SenderCompletion) String() string { _ = "STUB: not implemented"; return "" }
