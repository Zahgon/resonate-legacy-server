package t_aio

type Kind int

const (
	Echo Kind = iota
	Router
	Sender
	Store
)

func (k Kind) String() string { _ = "STUB: not implemented"; return "" }

type Submission struct {
	Kind Kind
	Tags map[string]string

	Echo   *EchoSubmission
	Router *RouterSubmission
	Sender *SenderSubmission
	Store  *StoreSubmission
}

func (s *Submission) String() string { _ = "STUB: not implemented"; return "" }

type Completion struct {
	Kind Kind
	Tags map[string]string

	Echo   *EchoCompletion
	Router *RouterCompletion
	Sender *SenderCompletion
	Store  *StoreCompletion
}

func (c *Completion) String() string { _ = "STUB: not implemented"; return "" }
