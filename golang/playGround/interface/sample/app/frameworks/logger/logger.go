package logger

import "../../usercases/port"

// 実装確認
var _ port.Logger = (*Logger)(nil)

type Logger struct{}

func NewLogger() port.Logger {
	return &Logger{}
}

func (l *Logger) Info(msg string) {
	// Implement info level logging
}

func (l *Logger) Error(msg string) {
	// Implement error level logging
}
