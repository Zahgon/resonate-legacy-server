package util

import (
	"cmp"
	"time"

	"github.com/robfig/cron/v3"
)

func Assert(cond bool, msg string) { _ = "STUB: not implemented"; return }

type KV[K any, V any] struct {
	Key   K
	Value V
}

func OrderedRange[K cmp.Ordered, V any](m map[K]V) []V { _ = "STUB: not implemented"; return nil }

func OrderedRangeKV[K cmp.Ordered, V any](m map[K]V) []*KV[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func orderedRangeSort[K cmp.Ordered, V any](m map[K]V) []K { _ = "STUB: not implemented"; return nil }

// nosemgrep: range-over-map

func ToPointer[T any](val T) *T { _ = "STUB: not implemented"; return nil }

func SafeDeref[T any](val *T) T { _ = "STUB: not implemented"; return *new(T) }

func Next(curr int64, cronExp string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func unixMilliToTime(unixMilli int64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func ParseCron(cronExp string) (cron.Schedule, error) {
	_ = "STUB: not implemented"
	return *new(cron.Schedule), nil
}

func UnmarshalChain(data []byte, vs ...any) error { _ = "STUB: not implemented"; return nil }

// reset v to zero value

func RemoveWhitespace(s string) string { _ = "STUB: not implemented"; return "" }

func DeferAndLog(f func() error) { _ = "STUB: not implemented"; return }

func InvokeId(promiseId string) string { _ = "STUB: not implemented"; return "" }

func ResumeId(rootPromiseId, promiseId string) string { _ = "STUB: not implemented"; return "" }

func NotifyId(promiseId, id string) string { _ = "STUB: not implemented"; return "" }

func ClampAddInt64(a, b int64) int64 { _ = "STUB: not implemented"; return 0 }
