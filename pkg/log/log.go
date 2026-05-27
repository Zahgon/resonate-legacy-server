package log

import (
	"log/slog"
)

const (
	DebugLevel = slog.LevelDebug
	InfoLevel  = slog.LevelInfo
	WarnLevel  = slog.LevelWarn
	ErrorLevel = slog.LevelError
	OffLevel   = slog.Level(1000)
)

// ParseLevel takes a string level and returns the slog log level constant.
func ParseLevel(lvl string) (slog.Level, error) {
	_ = "STUB: not implemented"
	return *new(slog.Level), nil
}
