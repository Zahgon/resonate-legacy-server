package task

import (
	"encoding/json"

	"github.com/resonatehq/resonate/pkg/message"
)

type FencingToken struct {
	TaskId      string
	TaskCounter int64
}

type Task struct {
	Id            string          `json:"id"`
	Counter       int             `json:"counter"`
	Timeout       int64           `json:"timeout"`
	RootPromiseId string          `json:"rootPromiseId"`
	ProcessId     *string         `json:"processId,omitempty"`
	State         State           `json:"-"`
	Recv          json.RawMessage `json:"-"`
	Mesg          *message.Mesg   `json:"-"`
	Attempt       int             `json:"-"`
	Ttl           int64           `json:"-"`
	ExpiresAt     int64           `json:"-"`
	CreatedOn     *int64          `json:"createdOn,omitempty"`
	CompletedOn   *int64          `json:"completedOn,omitempty"`
}

func (t *Task) String() string { _ = "STUB: not implemented"; return "" }

func (t1 *Task) Equals(t2 *Task) bool {
	_ = "STUB: not implemented"
	// for dst only
	return false
}

type State int

const (
	Init      State = 1 << iota // 1
	Enqueued                    // 2
	Claimed                     // 4
	Completed                   // 8
	Timedout                    // 16
)

func (s State) String() string { _ = "STUB: not implemented"; return "" }

func (s State) In(mask State) bool { _ = "STUB: not implemented"; return false }
