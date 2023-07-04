package logger

import "void-project/pkg"

type Level uint8

const (
	AllLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	SQLLevel
	ServerLevel
)

func (l Level) Name() string {
	return [...]string{"", "debug", "info", "warn", "error", "sql", "server"}[pkg.IfElse(l > ServerLevel, 0, l)]
}

func (l Level) Value() uint8 {
	return [...]uint8{0, 1, 2, 3, 4, 5, 6}[pkg.IfElse(l > ServerLevel, 0, l)]
}
