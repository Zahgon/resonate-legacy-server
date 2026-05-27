package http

import (
	"encoding/json"

	"github.com/resonatehq/resonate/pkg/promise"

	"github.com/gin-gonic/gin"
)

// Read

type readPromiseHeader struct {
	RequestId string `header:"request-id"`
}

func (s *server) readPromise(c *gin.Context) { _ = "STUB: not implemented"; return }

// Search

type searchPromisesHeader struct {
	RequestId string `header:"request-id"`
}

type searchPromisesParams struct {
	Id     *string           `form:"id" json:"id" binding:"omitempty,min=1"`
	State  *string           `form:"state" json:"state" binding:"omitempty,oneofcaseinsensitive=pending resolved rejected"`
	Tags   map[string]string `form:"tags" json:"tags,omitempty"`
	Limit  *int              `form:"limit" json:"limit" binding:"omitempty,gte=0,lte=100"`
	Cursor *string           `form:"cursor" json:"cursor,omitempty"`
}

func (s *server) searchPromises(c *gin.Context) { _ = "STUB: not implemented"; return }

// tags needs to be parsed manually
// see: https://github.com/gin-gonic/gin/issues/2606

// Create

type createPromiseHeader struct {
	RequestId   string `header:"request-id"`
	Traceparent string `header:"traceparent"`
	Tracestate  string `header:"tracestate"`
	TaskId      string `header:"task-id"`
	TaskCounter int64  `header:"task-counter"`
}

type createPromiseBody struct {
	Id      string            `json:"id" binding:"required"`
	Param   promise.Value     `json:"param"`
	Timeout int64             `json:"timeout"`
	Tags    map[string]string `json:"tags,omitempty"`
}

func (s *server) createPromise(c *gin.Context) { _ = "STUB: not implemented"; return }

type createPromiseAndTaskBody struct {
	Promise createPromiseBody     `json:"promise" binding:"required"`
	Task    createPromiseTaskBody `json:"task" binding:"required"`
}

type createPromiseTaskBody struct {
	ProcessId string `json:"processId" binding:"required"`
	Ttl       int64  `json:"ttl" binding:"min=0"`
}

func (s *server) createPromiseAndTask(c *gin.Context) { _ = "STUB: not implemented"; return }

// Complete

type completePromiseHeader struct {
	RequestId   string `header:"request-id"`
	TaskId      string `header:"task-id"`
	TaskCounter int64  `header:"task-counter"`
}

type completePromiseBody struct {
	State promise.State `json:"state" binding:"required"`
	Value promise.Value `json:"value"`
}

func (s *server) completePromise(c *gin.Context) { _ = "STUB: not implemented"; return }

// Callback

type createCallbackHeader struct {
	RequestId   string `header:"request-id"`
	Traceparent string `header:"traceparent"`
	Tracestate  string `header:"tracestate"`
	TaskId      string `header:"task-id"`
	TaskCounter int64  `header:"task-counter"`
}

type createCallbackBody struct {
	PromiseId     string          `json:"promiseId"`
	RootPromiseId string          `json:"rootPromiseId" binding:"required"`
	Recv          json.RawMessage `json:"recv" binding:"required"`
	Timeout       int64           `json:"timeout"`
}

func (s *server) createCallback(c *gin.Context) { _ = "STUB: not implemented"; return }

// The parameter takes priority, but once we remove the deprecated route we
// can remove promise id from the body and retrieve the information solely
// from the route parameter.

// Subscribe

type createSubscriptionHeader struct {
	RequestId   string `header:"request-id"`
	Traceparent string `header:"traceparent"`
	Tracestate  string `header:"tracestate"`
}

type createSubscriptionBody struct {
	Id        string          `json:"Id" binding:"required"`
	PromiseId string          `json:"promiseId"`
	Recv      json.RawMessage `json:"recv" binding:"required"`
	Timeout   int64           `json:"timeout"`
}

func (s *server) createSubscription(c *gin.Context) { _ = "STUB: not implemented"; return }

// The parameter takes priority, but once we remove the deprecated route we
// can remove promise id from the body and retrieve the information solely
// from the route parameter.
