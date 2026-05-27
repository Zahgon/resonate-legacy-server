package tasks

import (
	"github.com/resonatehq/resonate/pkg/client"
	"github.com/spf13/cobra"
)

// Example command usage for completing a task
var completeTasksExample = `
# Complete a task
resonate tasks complete foo --counter 1`

// CompleteTaskCmd returns a cobra command for completing a task.
func CompleteTaskCmd(c client.Client) *cobra.Command { _ = "STUB: not implemented"; return nil }

// Define the cobra command

// Validate required flags

// Create the body for the complete task request

// Call the client method to complete the task

// Return any errors from the request

// Handle the response based on the status code

// Return nil if no error occurred

// Define command flags

// Mark flags as required
