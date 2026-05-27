package promise

type PromiseRecord struct {
	Id           string
	State        State
	ParamHeaders []byte
	ParamData    []byte
	ValueHeaders []byte
	ValueData    []byte
	Timeout      int64
	CreatedOn    *int64
	CompletedOn  *int64
	Tags         []byte
	SortId       int64
}

func (r *PromiseRecord) Promise() (*Promise, error) { _ = "STUB: not implemented"; return nil, nil }

func bytesToMap(b []byte) (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }
