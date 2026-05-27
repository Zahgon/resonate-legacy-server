package t_api

type Kind int

const (
	_ Kind = iota
	// PROMISES
	PromiseGet
	PromiseSearch
	PromiseCreate
	PromiseComplete
	PromiseRegister

	// SCHEDULES
	ScheduleRead
	ScheduleSearch
	ScheduleCreate
	ScheduleDelete

	// TASKS
	TaskCreate
	TaskAcquire
	TaskRelease
	TaskComplete
	TaskHeartbeat

	// Echo
	Echo

	// Noop
	Noop
)

func (k Kind) String() string {
	_ = "STUB: not implemented"

	// PROMISES
	return ""
}

// SCHEDULES

// TASKS

// ECHO

// NOOP
