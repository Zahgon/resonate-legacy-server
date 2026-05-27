package schedules

import (
	"github.com/resonatehq/resonate/pkg/client"
	"github.com/spf13/cobra"
)

var createScheduleExample = `
# Create a schedule that runs every minute
resonate schedules create foo --cron "* * * * *" --promise-timeout 1h --promise-id "foo.{{.timestamp}}"

# Create a schedule that runs every 5 minutes and includes data and headers
resonate schedules create foo --cron "*/5 * * * *" --promise-timeout 1h --promise-id "foo.{{.timestamp}}" --promise-data foo --promise-header bar=bar`

func CreateScheduleCmd(c client.Client) *cobra.Command { _ = "STUB: not implemented"; return nil }
