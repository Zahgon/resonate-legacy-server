package dst

import (
	"math/rand"
)

type BcValidator struct {
	validators []BcValidatorFn
}

type BcValidatorFn func(*Model, int64, int64, *Req) (*Model, error)

func NewBcValidator(r *rand.Rand, config *Config) *BcValidator {
	_ = "STUB: not implemented"
	return nil
}

func (v *BcValidator) AddBcValidator(bcv BcValidatorFn) { _ = "STUB: not implemented"; return }

func (v *BcValidator) Validate(model *Model, reqTime int64, resTime int64, req *Req) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ValidateNotify(model *Model, reqTime int64, resTime int64, req *Req) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the only way this can happen is if the promise timedout

func ValidateTasksWithSameRootPromiseId(model *Model, reqTime int64, _ int64, req *Req) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ValidateTaskExpiry(model *Model, reqTime int64, _ int64, req *Req) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
