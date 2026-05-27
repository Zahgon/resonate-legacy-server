package message

type Mesg struct {
	Type Type              `json:"type"`
	Head map[string]string `json:"head,omitempty"`
	Root string            `json:"root"`
	Leaf string            `json:"leaf"`
}

func (m *Mesg) String() string { _ = "STUB: not implemented"; return "" }

type Type string

const (
	Invoke = "invoke"
	Resume = "resume"
	Notify = "notify"
)
