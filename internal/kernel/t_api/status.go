package t_api

// StatusCode represents the type of response that occurred
type StatusCode int

const (
	// Application level status (2000-4999)
	StatusOK        StatusCode = 20000
	StatusCreated   StatusCode = 20100
	StatusNoContent StatusCode = 20400

	StatusFieldValidationError   StatusCode = 40000
	StatusUnauthorized           StatusCode = 40100
	StatusForbidden              StatusCode = 40300
	StatusTaskAlreadyClaimed     StatusCode = 40306
	StatusTaskAlreadyCompleted   StatusCode = 40307
	StatusTaskInvalidCounter     StatusCode = 40308
	StatusTaskInvalidState       StatusCode = 40309
	StatusPromiseNotFound        StatusCode = 40400
	StatusScheduleNotFound       StatusCode = 40401
	StatusTaskNotFound           StatusCode = 40403
	StatusPromiseRecvNotFound    StatusCode = 40404
	StatusTaskPreconditionFailed StatusCode = 41200

	// Platform level status (50000-59909)
	StatusInternalServerError    StatusCode = 50000
	StatusAIOEchoError           StatusCode = 50001
	StatusAIOMatchError          StatusCode = 50002
	StatusAIOQueueError          StatusCode = 50003
	StatusAIOStoreError          StatusCode = 50004
	StatusSystemShuttingDown     StatusCode = 50300
	StatusAPISubmissionQueueFull StatusCode = 50301
	StatusAIOSubmissionQueueFull StatusCode = 50302
	StatusSchedulerQueueFull     StatusCode = 50303
)

// String returns the string representation of the status code.
func (s StatusCode) String() string { _ = "STUB: not implemented"; return "" }

func (s StatusCode) IsSuccessful() bool { _ = "STUB: not implemented"; return false }
