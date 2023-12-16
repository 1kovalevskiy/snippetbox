package logger

import (
	"log"
	"log/slog"
	"os"
)

type Interface interface {
	Debug(message string, args ...any)
	Info(message string, args ...any)
	Warn(message string, args ...any)
	Error(message string, args ...any)
	Printf(format string, args ...any)
	With(args ...any) *Logger
}

type Logger struct {
	slogger *slog.Logger
	Logger  *log.Logger
}

var _ Interface = (*Logger)(nil)

func NewLogger() *Logger {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	slogger := slog.New(handler)
	logger := slog.NewLogLogger(handler, slog.LevelInfo)
	return &Logger{slogger, logger}
}

func (l *Logger) Printf(format string, args ...any) {
	l.Logger.Printf(format, args...)
}

func (l *Logger) Fatal(args ...any) {
	l.Logger.Fatal(args...)
}

func (l *Logger) Debug(message string, args ...any) {
	l.slogger.Debug(message, args...)
}

func (l *Logger) Info(message string, args ...any) {
	l.slogger.Info(message, args...)
}

func (l *Logger) Warn(message string, args ...any) {
	l.slogger.Warn(message, args...)
}

func (l *Logger) Error(message string, args ...any) {
	l.slogger.Error(message, args...)
}

func (l *Logger) With(args ...any) *Logger {
	slogger := l.slogger.With(args...)
	return &Logger{slogger, l.Logger}
}
