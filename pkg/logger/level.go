package logger

import "void-project/pkg"

type Level uint8

const (
	RemoveAll Level = iota
	Debug
	Info
	Warn
	Error
	SQL
)

func (l Level) Name() string {
	return [...]string{"", "debug", "info", "warn", "error", "sql"}[pkg.IfElse(l > SQL, 0, l)]
}

func (l Level) Value() uint8 {
	return [...]uint8{0, 1, 2, 3, 4, 5}[pkg.IfElse(l > SQL, 0, l)]
}
