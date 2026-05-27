package tree

import (
	v1 "github.com/resonatehq/resonate/pkg/client/v1"
	"github.com/spf13/cobra"
)

type searchParams struct {
	id     string
	origin string
	cursor *string
}

func NewCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Try new tag format first (resonate:origin)

// If we got results with resonate:origin, we have the entire call graph

// Collect all promises with pagination

// Handle pagination

// Old format: global scope indicates a new branch

// nosemgrep: range-over-map

// Print the current node

// Recurse into children

// Flags

func details(p *v1.Promise) string { _ = "STUB: not implemented"; return "" }
