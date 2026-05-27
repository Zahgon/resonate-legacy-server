package t_api

import (
	"github.com/resonatehq/resonate/pkg/callback"
	"github.com/resonatehq/resonate/pkg/promise"
	"github.com/resonatehq/resonate/pkg/schedule"
	"github.com/resonatehq/resonate/pkg/task"
)

type Response struct {
	Status StatusCode
	Head   map[string]string
	Data   ResponsePayload
}

type ResponsePayload interface {
	Kind() Kind
	String() string
	isResponsePayload()
}

// Promises

type PromiseGetResponse struct {
	Promise *promise.Promise `json:"promise,omitempty"`
}

func (r *PromiseGetResponse) String() string { _ = "STUB: not implemented"; return "" }

func (r *PromiseGetResponse) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type PromiseSearchResponse struct {
	Promises []*promise.Promise            `json:"promises,omitempty"`
	Cursor   *Cursor[PromiseSearchRequest] `json:"cursor,omitempty"`
}

func (r *PromiseSearchResponse) String() string { _ = "STUB: not implemented"; return "" }

func (r *PromiseSearchResponse) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type PromiseCreateResponse struct {
	Promise *promise.Promise `json:"promise,omitempty"`
}

func (r *PromiseCreateResponse) String() string { _ = "STUB: not implemented"; return "" }

func (r *PromiseCreateResponse) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type TaskCreateResponse struct {
	Promise *promise.Promise `json:"promise,omitempty"`
	Task    *task.Task       `json:"task,omitempty"`
}

func (r *TaskCreateResponse) String() string { _ = "STUB: not implemented"; return "" }

func (r *TaskCreateResponse) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type PromiseCompleteResponse struct {
	Promise *promise.Promise `json:"promise,omitempty"`
}

func (r *PromiseCompleteResponse) String() string { _ = "STUB: not implemented"; return "" }

func (r *PromiseCompleteResponse) Kind() Kind {
	_ = "STUB: not implemented"
	return *

	// Callbacks
	new(Kind)
}

type PromiseRegisterResponse struct {
	Promise  *promise.Promise   `json:"promise,omitempty"`
	Callback *callback.Callback `json:"callback,omitempty"`
}

func (r *PromiseRegisterResponse) String() string { _ = "STUB: not implemented"; return "" }

func (r *PromiseRegisterResponse) Kind() Kind {
	_ = "STUB: not implemented"
	return *

	// Schedules
	new(Kind)
}

type ScheduleGetResponse struct {
	Schedule *schedule.Schedule `json:"schedule,omitempty"`
}

func (r *ScheduleGetResponse) String() string { _ = "STUB: not implemented"; return "" }

func (r *ScheduleGetResponse) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type ScheduleSearchResponse struct {
	Schedules []*schedule.Schedule           `json:"schedules,omitempty"`
	Cursor    *Cursor[ScheduleSearchRequest] `json:"cursor,omitempty"`
}

func (r *ScheduleSearchResponse) String() string { _ = "STUB: not implemented"; return "" }

func (r *ScheduleSearchResponse) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type ScheduleCreateResponse struct {
	Schedule *schedule.Schedule `json:"schedule,omitempty"`
}

func (r *ScheduleCreateResponse) String() string { _ = "STUB: not implemented"; return "" }

func (r *ScheduleCreateResponse) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type ScheduleDeleteResponse struct{}

func (r *ScheduleDeleteResponse) String() string { _ = "STUB: not implemented"; return "" }

func (r *ScheduleDeleteResponse) Kind() Kind {
	_ = "STUB: not implemented"
	return *

	// Tasks
	new(Kind)
}

type TaskAcquireResponse struct {
	Task            *task.Task       `json:"task,omitempty"`
	RootPromise     *promise.Promise `json:"rootPromise,omitempty"`
	LeafPromise     *promise.Promise `json:"leafPromise,omitempty"`
	RootPromiseHref string           `json:"rootPromiseHref,omitempty"`
	LeafPromiseHref string           `json:"leafPromiseHref,omitempty"`
}

func (r *TaskAcquireResponse) String() string { _ = "STUB: not implemented"; return "" }

func (r *TaskAcquireResponse) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type TaskCompleteResponse struct {
	Task *task.Task `json:"task,omitempty"`
}

func (r *TaskCompleteResponse) String() string { _ = "STUB: not implemented"; return "" }

func (r *TaskCompleteResponse) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type TaskReleaseResponse struct {
	Task *task.Task `json:"task,omitempty"`
}

func (r *TaskReleaseResponse) String() string { _ = "STUB: not implemented"; return "" }

func (r *TaskReleaseResponse) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

type TaskHeartbeatResponse struct {
	TasksAffected int64 `json:"tasksAffected"`
}

func (r *TaskHeartbeatResponse) String() string { _ = "STUB: not implemented"; return "" }

func (r *TaskHeartbeatResponse) Kind() Kind {
	_ = "STUB: not implemented"
	return *

	// Echo
	new(Kind)
}

type EchoResponse struct {
	Data string `json:"data"`
}

func (r *EchoResponse) String() string { _ = "STUB: not implemented"; return "" }

func (r *EchoResponse) Kind() Kind {
	_ = "STUB: not implemented"

	// Noop
	return *new(Kind)
}

type NoopResponse struct{}

func (r *NoopResponse) String() string { _ = "STUB: not implemented"; return "" }

func (r *NoopResponse) Kind() Kind {
	_ = "STUB: not implemented"

	// Marker methods that make each of the request types be a
	// ResponsePayload type.
	return *new(Kind)
}

func (r *PromiseGetResponse) isResponsePayload()      { _ = "STUB: not implemented"; return }
func (r *PromiseSearchResponse) isResponsePayload()   { _ = "STUB: not implemented"; return }
func (r *PromiseCreateResponse) isResponsePayload()   { _ = "STUB: not implemented"; return }
func (r *TaskCreateResponse) isResponsePayload()      { _ = "STUB: not implemented"; return }
func (r *PromiseCompleteResponse) isResponsePayload() { _ = "STUB: not implemented"; return }
func (r *PromiseRegisterResponse) isResponsePayload() { _ = "STUB: not implemented"; return }
func (r *ScheduleGetResponse) isResponsePayload()     { _ = "STUB: not implemented"; return }
func (r *ScheduleSearchResponse) isResponsePayload()  { _ = "STUB: not implemented"; return }
func (r *ScheduleCreateResponse) isResponsePayload()  { _ = "STUB: not implemented"; return }
func (r *ScheduleDeleteResponse) isResponsePayload()  { _ = "STUB: not implemented"; return }
func (r *TaskAcquireResponse) isResponsePayload()     { _ = "STUB: not implemented"; return }
func (r *TaskCompleteResponse) isResponsePayload()    { _ = "STUB: not implemented"; return }
func (r *TaskReleaseResponse) isResponsePayload()     { _ = "STUB: not implemented"; return }
func (r *TaskHeartbeatResponse) isResponsePayload()   { _ = "STUB: not implemented"; return }
func (r *EchoResponse) isResponsePayload()            { _ = "STUB: not implemented"; return }
func (r *NoopResponse) isResponsePayload()            { _ = "STUB: not implemented"; return }

func (r *Response) String() string { _ = "STUB: not implemented"; return "" }

func (r *Response) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

// Methods to cast Response.Payload to specific response payload types (direct assertion)
func (r *Response) AsPromiseGetResponse() *PromiseGetResponse {
	_ = "STUB: not implemented"
	return nil
}

func (r *Response) AsPromiseSearchResponse() *PromiseSearchResponse {
	_ = "STUB: not implemented"
	return nil
}

func (r *Response) AsPromiseCreateResponse() *PromiseCreateResponse {
	_ = "STUB: not implemented"
	return nil
}

func (r *Response) AsTaskCreateResponse() *TaskCreateResponse {
	_ = "STUB: not implemented"
	return nil
}

func (r *Response) AsPromiseCompleteResponse() *PromiseCompleteResponse {
	_ = "STUB: not implemented"
	return nil
}

func (r *Response) AsPromiseRegisterResponse() *PromiseRegisterResponse {
	_ = "STUB: not implemented"
	return nil
}

func (r *Response) AsScheduleGetResponse() *ScheduleGetResponse {
	_ = "STUB: not implemented"
	return nil
}

func (r *Response) AsScheduleSearchResponse() *ScheduleSearchResponse {
	_ = "STUB: not implemented"
	return nil
}

func (r *Response) AsScheduleCreateResponse() *ScheduleCreateResponse {
	_ = "STUB: not implemented"
	return nil
}

func (r *Response) AsScheduleDeleteResponse() *ScheduleDeleteResponse {
	_ = "STUB: not implemented"
	return nil
}

func (r *Response) AsTaskAcquireResponse() *TaskAcquireResponse {
	_ = "STUB: not implemented"
	return nil
}

func (r *Response) AsTaskCompleteResponse() *TaskCompleteResponse {
	_ = "STUB: not implemented"
	return nil
}

func (r *Response) AsTaskReleaseResponse() *TaskReleaseResponse {
	_ = "STUB: not implemented"
	return nil
}

func (r *Response) AsTaskHeartbeatResponse() *TaskHeartbeatResponse {
	_ = "STUB: not implemented"
	return nil
}

func (r *Response) AsEchoResponse() *EchoResponse { _ = "STUB: not implemented"; return nil }
