package slog

import (
	"log/slog"
	"os"
)

var SLog *slog.Logger

func init() {
	SLog = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func Debug(msg string, args ...any) {
	SLog.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	SLog.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	SLog.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	SLog.Error(msg, args...)
}
