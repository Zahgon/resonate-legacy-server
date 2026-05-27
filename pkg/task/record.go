package task

type TaskRecord struct {
	Id            string
	ProcessId     *string
	State         State
	RootPromiseId string
	Recv          []byte
	Mesg          []byte
	Timeout       int64
	Counter       int
	Attempt       int
	Ttl           int64
	ExpiresAt     int64
	CreatedOn     *int64
	CompletedOn   *int64
}

func (r *TaskRecord) Task() (*Task, error) { _ = "STUB: not implemented"; return nil, nil }
