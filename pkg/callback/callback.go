package callback

import (
	"encoding/json"

	"github.com/resonatehq/resonate/pkg/message"
)

type Callback struct {
	Id            string          `json:"id"`
	PromiseId     string          `json:"promiseId"`
	RootPromiseId string          `json:"rootPromiseId"`
	Recv          json.RawMessage `json:"-"`
	Mesg          *message.Mesg   `json:"-"`
	Timeout       int64           `json:"timeout"`
	CreatedOn     int64           `json:"createdOn"`
}

func (c *Callback) String() string { _ = "STUB: not implemented"; return "" }

func (c1 *Callback) Equals(c2 *Callback) bool {
	_ = "STUB: not implemented"
	// for dst only
	return false
}
