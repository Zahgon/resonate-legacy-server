package t_api

import (
	"encoding/json"

	"github.com/resonatehq/resonate/pkg/message"
	"github.com/resonatehq/resonate/pkg/promise"
)

type Request struct {
	Head map[string]string
	Data RequestPayload
}

type RequestPayload interface {
	Kind() Kind
	String() string
	Validate() error
	isRequestPayload()
}

// Promises

type PromiseGetRequest struct {
	Id string `json:"id"`
}

func (r *PromiseGetRequest) String() string { _ = "STUB: not implemented"; return "" }

func (r *PromiseGetRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (r *PromiseGetRequest) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type PromiseSearchRequest struct {
	Id     string            `json:"id"`
	States []promise.State   `json:"states"`
	Tags   map[string]string `json:"tags"`
	Limit  int               `json:"limit"`
	SortId *int64            `json:"sortId"`
}

func (r *PromiseSearchRequest) String() string { _ = "STUB: not implemented"; return "" }

func (r *PromiseSearchRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (r *PromiseSearchRequest) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type PromiseCreateRequest struct {
	Id      string            `json:"id"`
	Param   promise.Value     `json:"param,omitempty"`
	Timeout int64             `json:"timeout"`
	Tags    map[string]string `json:"tags,omitempty"`
}

func (r *PromiseCreateRequest) String() string { _ = "STUB: not implemented"; return "" }

func (r *PromiseCreateRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (r *PromiseCreateRequest) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type TaskCreateRequest struct {
	Promise *PromiseCreateRequest
	Task    *CreateTaskRequest
}

func (r *TaskCreateRequest) String() string { _ = "STUB: not implemented"; return "" }

func (r *TaskCreateRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (r *TaskCreateRequest) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type PromiseCompleteRequest struct {
	Id    string        `json:"id"`
	State promise.State `json:"state"`
	Value promise.Value `json:"value,omitempty"`
}

func (r *PromiseCompleteRequest) String() string { _ = "STUB: not implemented"; return "" }

func (r *PromiseCompleteRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (r *PromiseCompleteRequest) Kind() Kind {
	_ = "STUB: not implemented"
	return *

	// Callbacks
	new(Kind)
}

type PromiseRegisterRequest struct {
	Id        string          `json:"id"`
	PromiseId string          `json:"promiseId"`
	Recv      json.RawMessage `json:"recv"`
	Mesg      *message.Mesg   `json:"mesg"`
	Timeout   int64           `json:"timeout"`
}

func (r *PromiseRegisterRequest) String() string { _ = "STUB: not implemented"; return "" }

func (r *PromiseRegisterRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (r *PromiseRegisterRequest) Kind() Kind {
	_ = "STUB: not implemented"
	return *

	// Schedules
	new(Kind)
}

type ScheduleGetRequest struct {
	Id string `json:"id"`
}

func (r *ScheduleGetRequest) String() string { _ = "STUB: not implemented"; return "" }

func (r *ScheduleGetRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (r *ScheduleGetRequest) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type ScheduleSearchRequest struct {
	Id     string            `json:"id"`
	Tags   map[string]string `json:"tags"`
	Limit  int               `json:"limit"`
	SortId *int64            `json:"sortId"`
}

func (r *ScheduleSearchRequest) String() string { _ = "STUB: not implemented"; return "" }

func (r *ScheduleSearchRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (r *ScheduleSearchRequest) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type ScheduleCreateRequest struct {
	Id             string            `json:"id"`
	Description    string            `json:"desc,omitempty"`
	Cron           string            `json:"cron"`
	Tags           map[string]string `json:"tags,omitempty"`
	PromiseId      string            `json:"promiseId"`
	PromiseTimeout int64             `json:"promiseTimeout"`
	PromiseParam   promise.Value     `json:"promiseParam,omitempty"`
	PromiseTags    map[string]string `json:"promiseTags,omitempty"`
}

func (r *ScheduleCreateRequest) String() string { _ = "STUB: not implemented"; return "" }

func (r *ScheduleCreateRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (r *ScheduleCreateRequest) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type ScheduleDeleteRequest struct {
	Id string `json:"id"`
}

func (r *ScheduleDeleteRequest) String() string { _ = "STUB: not implemented"; return "" }

func (r *ScheduleDeleteRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (r *ScheduleDeleteRequest) Kind() Kind {
	_ = "STUB: not implemented"
	return *

	// Tasks
	new(Kind)
}

// CreateTaskRequest is not a Request on its own it needs to be part of a CreatePromiseAndTask Request
// that is why it does not implement the request interface
type CreateTaskRequest struct {
	PromiseId string `json:"promiseId"`
	ProcessId string `json:"processId"`
	Ttl       int64  `json:"ttl" binding:"min=0"`
	Timeout   int64  `json:"timeout"`
}

func (r *CreateTaskRequest) String() string { _ = "STUB: not implemented"; return "" }

type TaskAcquireRequest struct {
	Id        string `json:"id"`
	Counter   int    `json:"counter"`
	ProcessId string `json:"processId"`
	Ttl       int64  `json:"ttl" binding:"min=0"`
}

func (r *TaskAcquireRequest) String() string { _ = "STUB: not implemented"; return "" }

func (r *TaskAcquireRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (r *TaskAcquireRequest) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type TaskCompleteRequest struct {
	Id      string `json:"id"`
	Counter int    `json:"counter"`
}

func (r *TaskCompleteRequest) String() string { _ = "STUB: not implemented"; return "" }

func (r *TaskCompleteRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (r *TaskCompleteRequest) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type TaskReleaseRequest struct {
	Id      string `json:"id"`
	Counter int    `json:"counter"`
}

func (r *TaskReleaseRequest) String() string { _ = "STUB: not implemented"; return "" }

func (r *TaskReleaseRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (r *TaskReleaseRequest) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type TaskHeartbeatRequest struct {
	ProcessId string `json:"processId"`
}

func (r *TaskHeartbeatRequest) String() string { _ = "STUB: not implemented"; return "" }

func (r *TaskHeartbeatRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (r *TaskHeartbeatRequest) Kind() Kind {
	_ = "STUB: not implemented"
	return *

	// Echo
	new(Kind)
}

type EchoRequest struct {
	Data string `json:"data"`
}

func (r *EchoRequest) String() string { _ = "STUB: not implemented"; return "" }

func (r *EchoRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (r *EchoRequest) Kind() Kind {
	_ = "STUB: not implemented"

	// Noop
	return *new(Kind)
}

type NoopRequest struct{}

func (r *NoopRequest) String() string { _ = "STUB: not implemented"; return "" }

func (r *NoopRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (r *NoopRequest) Kind() Kind {
	_ = "STUB: not implemented"

	// Marker methods that make each of the request types be a
	// RequestPayload type.
	return *new(Kind)
}

func (r *PromiseGetRequest) isRequestPayload()      { _ = "STUB: not implemented"; return }
func (r *PromiseSearchRequest) isRequestPayload()   { _ = "STUB: not implemented"; return }
func (r *PromiseCreateRequest) isRequestPayload()   { _ = "STUB: not implemented"; return }
func (r *TaskCreateRequest) isRequestPayload()      { _ = "STUB: not implemented"; return }
func (r *PromiseCompleteRequest) isRequestPayload() { _ = "STUB: not implemented"; return }
func (r *PromiseRegisterRequest) isRequestPayload() { _ = "STUB: not implemented"; return }
func (r *ScheduleGetRequest) isRequestPayload()     { _ = "STUB: not implemented"; return }
func (r *ScheduleSearchRequest) isRequestPayload()  { _ = "STUB: not implemented"; return }
func (r *ScheduleCreateRequest) isRequestPayload()  { _ = "STUB: not implemented"; return }
func (r *ScheduleDeleteRequest) isRequestPayload()  { _ = "STUB: not implemented"; return }
func (r *TaskAcquireRequest) isRequestPayload()     { _ = "STUB: not implemented"; return }
func (r *TaskCompleteRequest) isRequestPayload()    { _ = "STUB: not implemented"; return }
func (r *TaskReleaseRequest) isRequestPayload()     { _ = "STUB: not implemented"; return }
func (r *TaskHeartbeatRequest) isRequestPayload()   { _ = "STUB: not implemented"; return }
func (r *EchoRequest) isRequestPayload()            { _ = "STUB: not implemented"; return }
func (r *NoopRequest) isRequestPayload() {
	_ = "STUB: not implemented"

	// Request Methods
	return
}

func (r *Request) String() string { _ = "STUB: not implemented"; return "" }

func (r *Request) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

func (r *Request) Validate() error { _ = "STUB: not implemented"; return nil }
