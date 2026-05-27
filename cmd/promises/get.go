package promises

import (
	"github.com/resonatehq/resonate/pkg/client"
	"github.com/spf13/cobra"
)

var getPromiseExample = `
# Get a promise
resonate promises foo`

func GetPromiseCmd(c client.Client) *cobra.Command { _ = "STUB: not implemented"; return nil }
