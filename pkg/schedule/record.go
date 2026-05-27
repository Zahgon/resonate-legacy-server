package schedule

type ScheduleRecord struct {
	Id                  string
	Description         string
	Cron                string
	Tags                []byte
	PromiseId           string
	PromiseTimeout      int64
	PromiseParamHeaders []byte
	PromiseParamData    []byte
	PromiseTags         []byte
	LastRunTime         *int64
	NextRunTime         int64
	CreatedOn           int64
	SortId              int64
}

func (r *ScheduleRecord) Schedule() (*Schedule, error) { _ = "STUB: not implemented"; return nil, nil }

func bytesToMap(b []byte) (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }
