package promises

import (
	"github.com/resonatehq/resonate/pkg/client"
	"github.com/spf13/cobra"
)

var createPromiseExample = `
# Create a promise
resonate promises create foo --timeout 1h

# Create a promise with data and headers and tags
resonate promises create foo --timeout 1h --data foo --header bar=bar --tag baz=baz`

func CreatePromiseCmd(c client.Client) *cobra.Command { _ = "STUB: not implemented"; return nil }
