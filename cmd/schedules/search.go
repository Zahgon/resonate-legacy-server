package schedules

import (
	"github.com/resonatehq/resonate/pkg/client"
	"github.com/spf13/cobra"
)

var searchSchedulesExample = `
# Search for all schedules
resonate schedules search "*"

# Search for schedules that start with foo
resonate schedules search "foo.*"`

func SearchSchedulesCmd(c client.Client) *cobra.Command { _ = "STUB: not implemented"; return nil }
