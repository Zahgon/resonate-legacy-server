package promise

type Promise struct {
	Id          string            `json:"id"`
	State       State             `json:"state"`
	Param       Value             `json:"param,omitempty"`
	Value       Value             `json:"value,omitempty"`
	Timeout     int64             `json:"timeout"`
	Tags        map[string]string `json:"tags,omitempty"`
	CreatedOn   *int64            `json:"createdOn,omitempty"`
	CompletedOn *int64            `json:"completedOn,omitempty"`
	SortId      int64             `json:"-"` // unexported
}

func (p *Promise) String() string { _ = "STUB: not implemented"; return "" }

func GetTimedoutState(tags map[string]string) State { _ = "STUB: not implemented"; return *new(State) }

func (p1 *Promise) Equals(p2 *Promise) bool {
	_ = "STUB: not implemented"
	// for dst only
	return false
}

type State int

const (
	Pending  State = 1 << iota // 1
	Resolved                   // 2
	Rejected                   // 4
	Canceled                   // 8
	Timedout                   // 16
)

func (s State) String() string { _ = "STUB: not implemented"; return "" }

func (s *State) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *State) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s State) In(mask State) bool { _ = "STUB: not implemented"; return false }

type Value struct {
	Headers map[string]string `json:"headers,omitempty"`
	Data    []byte            `json:"data,omitempty"`
}

func (v Value) String() string { _ = "STUB: not implemented"; return "" }
