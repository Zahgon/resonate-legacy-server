package schedules

import (
	"github.com/resonatehq/resonate/pkg/client"
	"github.com/spf13/cobra"
)

var deleteScheduleExample = `
# Delete a schedule
resonate schedules delete foo`

func DeleteScheduleCmd(c client.Client) *cobra.Command { _ = "STUB: not implemented"; return nil }
