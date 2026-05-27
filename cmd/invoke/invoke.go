package invoke

import (
	"github.com/spf13/cobra"
)

var invokeExample = `
# Invoke a function with arguments
resonate invoke promise-id --func add --arg 1 --arg 2 --arg "a string"

# Invoke with timeout and target
resonate invoke promise-id --func process --arg data1 --arg 5 --timeout 1h --target "poll://any@default"

# Invoke with JSON arguments
resonate invoke promise-id --func process --json-args '[{"key": "value"}, {"num": 42}]'

# Invoke with version
resonate invoke promise-id --func process --version 2 --arg data`

type Param struct {
	Func    string `json:"func"`
	Args    []any  `json:"args"`
	Version int    `json:"version"`
}

func NewCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// If unmarshal fails, treat as string
