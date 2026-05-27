package promises

import (
	"github.com/resonatehq/resonate/pkg/client"
	"github.com/spf13/cobra"
)

var searchPromisesExample = `
# Search for all promises
resonate promises search "*"

# Search for promises that start with foo
resonate promises search "foo.*"

# Search for all pending promises
resonate promises search "*" --state pending

# Search for all resolved promises
resonate promises search "*" --state resolved

# Search for all rejected promises
resonate promises search "*" --state rejected`

func SearchPromisesCmd(c client.Client) *cobra.Command { _ = "STUB: not implemented"; return nil }
