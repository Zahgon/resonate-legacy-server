package dst

import (
	"github.com/spf13/cobra"
)

type Issue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

const issueFmt = `# DST Failed
%s

**Seed**
~~~
%d
~~~

**Scenario**
~~~
%s
~~~

**Store**
~~~
%s
~~~

**Commit**
~~~
%s
~~~

**Command**
~~~
go run ./... dst run --seed %d --ticks %d --scenario %s --aio-store-%s-enable
~~~

**Logs**
~~~
%s
~~~

[more details](%s)
`

func CreateDSTIssueCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// read logs file

// create github issue

func parseLogs(filename string, head int, tail int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func createGitHubIssue(repo string, token string, issue *Issue) error {
	_ = "STUB: not implemented"
	return nil
}

// Convert issue to JSON

// Create HTTP request

// Set required headers

// Send the HTTP request

// Check the response status
