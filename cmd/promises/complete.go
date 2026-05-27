package promises

import (
	"github.com/resonatehq/resonate/pkg/client"
	"github.com/spf13/cobra"
)

var completePromiseExample = `
# %s a promise
resonate promises %s foo

# %s a promise with data and headers
resonate promises %s foo --data foo --header bar=bar`

func CompletePromiseCmds(c client.Client) []*cobra.Command { _ = "STUB: not implemented"; return nil }
