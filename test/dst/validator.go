package dst

import (
	"math/rand"
	"regexp"

	"github.com/resonatehq/resonate/internal/kernel/t_api"
)

type Validator struct {
	regexes    map[string]*regexp.Regexp
	validators map[t_api.Kind]ResponseValidator
}

type ResponseValidator func(*Model, int64, int64, *t_api.Request, *t_api.Response) (*Model, error)

func NewValidator(r *rand.Rand, config *Config) *Validator { _ = "STUB: not implemented"; return nil }

func (v *Validator) AddValidator(kind t_api.Kind, validator ResponseValidator) {
	_ = "STUB: not implemented"
	return
}

func (v *Validator) Validate(model *Model, reqTime int64, resTime int64, req *t_api.Request, res *t_api.Response) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PROMISES

func (v *Validator) ValidateReadPromise(model *Model, reqTime int64, resTime int64, req *t_api.Request, res *t_api.Response) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the only way this can happen is if the promise timedout

func (v *Validator) ValidateSearchPromises(model *Model, reqTime int64, resTime int64, req *t_api.Request, res *t_api.Response) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ignore scheduled promises

func (v *Validator) ValidateCreatePromise(model *Model, reqTime int64, resTime int64, req *t_api.Request, res *t_api.Response) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *Validator) ValidateCreatePromiseAndTask(model *Model, reqTime int64, resTime int64, req *t_api.Request, res *t_api.Response) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *Validator) validateCreatePromise(model *Model, reqTime int64, resTime int64, req *t_api.PromiseCreateRequest, status t_api.StatusCode, res *t_api.PromiseCreateResponse) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// update model state

// the only way this can happen with this test setup is if the promise timedout

func (v *Validator) ValidateCompletePromise(model *Model, reqTime int64, resTime int64, req *t_api.Request, res *t_api.Response) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// update model state

// the only way this can happen is if the promise timedout

func (v *Validator) ValidateCreateCallback(model *Model, reqTime int64, resTime int64, req *t_api.Request, res *t_api.Response) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the status is Ok there has to be a promise

// If the promise is completed we don't create a callback but return 200

// If the promise is pending is possible that it has been timeout so we update the promises model
// if the promise is timedout we can handle it as if it was completed

// otherwise verify the callback was created previously

func (v *Validator) ValidateReadSchedule(model *Model, reqTime int64, resTime int64, req *t_api.Request, res *t_api.Response) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *Validator) ValidateSearchSchedules(model *Model, reqTime int64, resTime int64, req *t_api.Request, res *t_api.Response) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *Validator) ValidateCreateSchedule(model *Model, reqTime int64, resTime int64, req *t_api.Request, res *t_api.Response) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *Validator) ValidateDeleteSchedule(model *Model, reqTime int64, resTime int64, req *t_api.Request, res *t_api.Response) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TASKS

func (v *Validator) ValidateClaimTask(model *Model, reqTime int64, resTime int64, req *t_api.Request, res *t_api.Response) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *Validator) ValidateCompleteTask(model *Model, reqTime int64, resTime int64, req *t_api.Request, res *t_api.Response) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the response should have been status created if:
// - the task was not already completed or timedout
// - the task has not timedout
// - the root promise is still pending

func (v *Validator) ValidateDropTask(model *Model, reqTime int64, resTime int64, req *t_api.Request, res *t_api.Response) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the response should have been StatusCreated if:
// - the task was claimed
// - the task has not expired or timedout
// - the root promise is still pending

func (v *Validator) ValidateHeartbeatTasks(model *Model, reqTime int64, resTime int64, req *t_api.Request, res *t_api.Response) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// copy the model only once

// we can only update the model for tasks that are unambiguously
// heartbeated, and it's only an approximation
