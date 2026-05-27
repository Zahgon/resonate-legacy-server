package http

import (
	"github.com/gin-gonic/gin"
	"github.com/resonatehq/resonate/pkg/promise"
)

// Read

type readScheduleHeader struct {
	RequestId string `header:"request-id"`
}

func (s *server) readSchedule(c *gin.Context) { _ = "STUB: not implemented"; return }

// Search

type searchSchedulesHeader struct {
	RequestId string `header:"request-id"`
}

type searchSchedulesParams struct {
	Id     *string           `form:"id" json:"id,omitempty" binding:"omitempty,min=1"`
	Tags   map[string]string `form:"tags" json:"tags,omitempty"`
	Limit  *int              `form:"limit" json:"limit,omitempty" binding:"omitempty,gte=0,lte=100"`
	Cursor *string           `form:"cursor" json:"cursor,omitempty"`
}

func (s *server) searchSchedules(c *gin.Context) { _ = "STUB: not implemented"; return }

// tags needs to be parsed manually
// see: https://github.com/gin-gonic/gin/issues/2606

// Create

type createScheduleHeader struct {
	RequestId string `header:"request-id"`
}

type createScheduleBody struct {
	Id             string            `json:"id" binding:"required"`
	Description    string            `json:"desc,omitempty"`
	Cron           string            `json:"cron" binding:"required"`
	Tags           map[string]string `json:"tags,omitempty"`
	PromiseId      string            `json:"promiseId" binding:"required"`
	PromiseTimeout int64             `json:"promiseTimeout"`
	PromiseParam   promise.Value     `json:"promiseParam,omitempty"`
	PromiseTags    map[string]string `json:"promiseTags,omitempty"`
}

func (s *server) createSchedule(c *gin.Context) { _ = "STUB: not implemented"; return }

// Delete

type deleteScheduleHeader struct {
	RequestId string `header:"request-id"`
}

func (s *server) deleteSchedule(c *gin.Context) { _ = "STUB: not implemented"; return }

// Serves as a type assertion
