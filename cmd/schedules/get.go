package schedules

import (
	"github.com/resonatehq/resonate/pkg/client"
	"github.com/spf13/cobra"
)

var getScheduleExample = `
# Get a schedule
resonate schedules get foo`

func GetScheduleCmd(c client.Client) *cobra.Command { _ = "STUB: not implemented"; return nil }
