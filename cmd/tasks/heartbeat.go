package tasks

import (
	"github.com/resonatehq/resonate/pkg/client"
	"github.com/spf13/cobra"
)

// Example command usage for sending a heartbeat to a task
var heartbeatTasksExample = `
# Heartbeat a task
resonate tasks heartbeat --pid bar`

// HeartbeatTaskCmd returns a cobra command for sending a heartbeat to a task.
func HeartbeatTaskCmd(c client.Client) *cobra.Command { _ = "STUB: not implemented"; return nil }

// Define the cobra command

// Create parameters for the request

// Create the body for heartbeat tasks request

// Call the client method to send the heartbeat (GET request with path params)

// Return any errors from the request

// Handle the response based on the status code

// Return nil if no error occurred

// Define command flags
