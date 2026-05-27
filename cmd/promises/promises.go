package promises

import (
	v1 "github.com/resonatehq/resonate/pkg/client/v1"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Add subcommands

// Flags

func prettyPrintPromises(cmd *cobra.Command, promises ...v1.Promise) {
	_ = "STUB: not implemented"
	return
}

func prettyPrintPromise(cmd *cobra.Command, promise *v1.Promise) { _ = "STUB: not implemented"; return }
