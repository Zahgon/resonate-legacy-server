package t_aio

type EchoSubmission struct {
	Data string
}

func (s *EchoSubmission) String() string { _ = "STUB: not implemented"; return "" }

type EchoCompletion struct {
	Data string
}

func (c *EchoCompletion) String() string { _ = "STUB: not implemented"; return "" }
