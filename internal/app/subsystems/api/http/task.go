package http

import (
	"github.com/gin-gonic/gin"
)

// Claim

type claimTaskHeader struct {
	RequestId string `header:"request-id"`
}

type claimTaskBody struct {
	Id        string `json:"id" binding:"required"`
	Counter   int    `json:"counter" binding:"required"`
	ProcessId string `json:"processId" binding:"required"`
	Ttl       int64  `json:"ttl" binding:"min=0"`
}

func (s *server) claimTask(c *gin.Context) { _ = "STUB: not implemented"; return }

// Complete

type completeTaskHeader struct {
	RequestId string `header:"request-id"`
}

type completeTaskBody struct {
	Id      string `json:"id" binding:"required"`
	Counter int    `json:"counter" binding:"required"`
}

func (s *server) completeTask(c *gin.Context) { _ = "STUB: not implemented"; return }

// Drop tasks
type dropTaskHeader struct {
	RequestId string `header:"request-id"`
}

type dropTaskBody struct {
	Id      string `json:"id" binding:"required"`
	Counter int    `json:"counter" binding:"required"`
}

func (s *server) dropTask(c *gin.Context) { _ = "STUB: not implemented"; return }

// Serves as a type assertion

// Heartbeat

type heartbeatTasksHeader struct {
	RequestId string `header:"request-id"`
}

type heartbeatTaskBody struct {
	ProcessId string `json:"processId" binding:"required"`
}

func (s *server) heartbeatTasks(c *gin.Context) { _ = "STUB: not implemented"; return }
