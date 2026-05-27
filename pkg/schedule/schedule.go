package schedule

import (
	"github.com/resonatehq/resonate/pkg/promise"
)

type Schedule struct {
	Id             string            `json:"id"`
	Description    string            `json:"desc,omitempty"`
	Cron           string            `json:"cron"`
	Tags           map[string]string `json:"tags"`
	PromiseId      string            `json:"promiseId"`
	PromiseTimeout int64             `json:"promiseTimeout"`
	PromiseParam   promise.Value     `json:"promiseParam,omitempty"`
	PromiseTags    map[string]string `json:"promiseTags,omitempty"`
	LastRunTime    *int64            `json:"lastRunTime,omitempty"`
	NextRunTime    int64             `json:"nextRunTime"`
	CreatedOn      int64             `json:"createdOn"`
	SortId         int64             `json:"-"` // unexported
}

func (s *Schedule) String() string { _ = "STUB: not implemented"; return "" }

func (s1 *Schedule) Equals(s2 *Schedule) bool {
	_ = "STUB: not implemented"
	// for dst only
	return false
}
