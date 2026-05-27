package dst

import (
	"github.com/resonatehq/resonate/pkg/callback"
	"github.com/resonatehq/resonate/pkg/promise"
	"github.com/resonatehq/resonate/pkg/schedule"
	"github.com/resonatehq/resonate/pkg/task"
)

// Model

type Model struct {
	promises  *Store[string, *promise.Promise]
	callbacks *Store[string, *callback.Callback]
	schedules *Store[string, *schedule.Schedule]
	tasks     *Store[string, *task.Task]
}

func NewModel() *Model { _ = "STUB: not implemented"; return nil }

func (m *Model) Copy() *Model { _ = "STUB: not implemented"; return nil }

func (m1 *Model) Equals(m2 *Model) bool { _ = "STUB: not implemented"; return false }

// Store

type relatable interface {
	~int | ~int64 | ~string
}

type equatable[T any] interface {
	Equals(o T) bool
}

type Store[I relatable, T equatable[T]] []*struct {
	id    I
	value T
}

func (s *Store[I, T]) copy() *Store[I, T] { _ = "STUB: not implemented"; return nil }

func (s *Store[I, T]) all() []T { _ = "STUB: not implemented"; return nil }

func (s *Store[I, T]) get(id I) T { _ = "STUB: not implemented"; return *new(T) }

func (s *Store[I, T]) set(id I, value T) { _ = "STUB: not implemented"; return }

func (s *Store[I, T]) delete(id I) { _ = "STUB: not implemented"; return }

func (s1 *Store[I, T]) equals(s2 *Store[I, T]) bool { _ = "STUB: not implemented"; return false }
