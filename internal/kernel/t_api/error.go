package t_api

type Error struct {
	code          StatusCode
	originalError error
}

func NewError(code StatusCode, error error) *Error { _ = "STUB: not implemented"; return nil }

func (e *Error) Code() StatusCode { _ = "STUB: not implemented"; return *new(StatusCode) }

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Error) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *Error) Is(target error) bool { _ = "STUB: not implemented"; return false }
