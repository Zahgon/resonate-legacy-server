package promises

import (
	"github.com/resonatehq/resonate/pkg/client"
	"github.com/spf13/cobra"
)

var createPromiseSubscriptionExample = `
# Create a subscription
resonate promises subscribe foo --id bar --timeout 1h --recv default

# Create a subscription with url
resonate promises subscribe foo --id bar --timeout 1h --recv poll://default/1

# Create a subscription with object
resonate promises subscribe foo --id bar --timeout 1h --recv {"type": "poll", "data": {"group": "default", "id": "2"}}
`

func CreateSubscriptionCmd(c client.Client) *cobra.Command { _ = "STUB: not implemented"; return nil }
