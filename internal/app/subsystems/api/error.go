package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/resonatehq/resonate/internal/kernel/t_api"
)

type Error struct {
	// Code is the internal code that indicates the type of error
	Code t_api.StatusCode `json:"code,omitempty"`

	// Message is the error message
	Message string `json:"message,omitempty"`

	// Details is a list of details about the error
	Details []*ErrorDetails `json:"details,omitempty"`
}

type ErrorDetails struct {
	// Type is the specific error type
	Type string `json:"@type,omitempty"`

	// Message is a human readable description of the error
	Message string `json:"message,omitempty"`

	// Domain is the domain of the error
	Domain string `json:"domain,omitempty"`

	// Metadata is additional information about the error
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func ServerError(err error) *Error { _ = "STUB: not implemented"; return nil }

func RequestError(status t_api.StatusCode) *Error { _ = "STUB: not implemented"; return nil }

func RequestValidationError(err error) *Error { _ = "STUB: not implemented"; return nil }

// Helper functions

func parseBindingError(errs ...error) []string { _ = "STUB: not implemented"; return nil }

func parseFieldError(e validator.FieldError) string { _ = "STUB: not implemented"; return "" }
