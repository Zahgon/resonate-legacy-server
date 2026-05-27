package util

import (
	"math/rand" // nosemgrep
	"time"

	"github.com/mitchellh/mapstructure"
)

type _range interface {
	int | int64 | float64 | time.Duration
}

type rangeFlag[T _range] struct {
	min     T
	max     T
	format  func(T) string
	parse   func(string) (T, error)
	resolve func(*rangeFlag[T], *rand.Rand) T
}

func NewRangeIntFlag(min int, max int) *rangeFlag[int] { _ = "STUB: not implemented"; return nil }

func NewRangeInt64Flag(min int64, max int64) *rangeFlag[int64] {
	_ = "STUB: not implemented"
	return nil
}

func NewRangeFloat64Flag(min float64, max float64) *rangeFlag[float64] {
	_ = "STUB: not implemented"
	return nil
}

func NewRangeDurationFlag(min time.Duration, max time.Duration) *rangeFlag[time.Duration] {
	_ = "STUB: not implemented"
	return nil
}

func (f *rangeFlag[T]) String() string { _ = "STUB: not implemented"; return "" }

func (f *rangeFlag[T]) Type() string { _ = "STUB: not implemented"; return "" }

func (f *rangeFlag[T]) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (f *rangeFlag[T]) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (f *rangeFlag[T]) Min() T { _ = "STUB: not implemented"; return *new(T) }

func (f *rangeFlag[T]) Max() T { _ = "STUB: not implemented"; return *new(T) }

func (f *rangeFlag[T]) Resolve(r *rand.Rand) T { _ = "STUB: not implemented"; return *new(T) }

// Helper functions

func StringToRange(r *rand.Rand) mapstructure.DecodeHookFunc {
	_ = "STUB: not implemented"
	return *new(mapstructure.DecodeHookFunc)
}
