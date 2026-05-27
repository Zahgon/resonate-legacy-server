package schedules

import (
	v1 "github.com/resonatehq/resonate/pkg/client/v1"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Add subcommands

// Flags

func prettyPrintSchedules(cmd *cobra.Command, schedules ...v1.Schedule) {
	_ = "STUB: not implemented"
	return
}

func prettyPrintSchedule(cmd *cobra.Command, schedule *v1.Schedule) {
	_ = "STUB: not implemented"
	return
}
