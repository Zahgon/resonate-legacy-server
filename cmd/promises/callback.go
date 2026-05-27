package promises

import (
	"github.com/resonatehq/resonate/pkg/client"
	"github.com/spf13/cobra"
)

var createPromiseCallbackExample = `
# Create a callback
resonate promises callback foo --root-promise-id bar --timeout 1h --recv default

# Create a callback with url
resonate promises callback foo --root-promise-id bar --timeout 1h --recv poll://default/1

# Create a callback with object
resonate promises callback foo --root-promise-id bar --timeout 1h --recv {"type": "poll", "data": {"group": "default", "id": "2"}}
`

func CreatePromiseCallbackCmd(c client.Client) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}
