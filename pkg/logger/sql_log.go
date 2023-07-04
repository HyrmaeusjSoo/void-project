package logger

import (
	"io"
	"log"
	"os"
)

type SQLLogger struct {
	lv Level
}

func NewSQLLogger() *SQLLogger {
	return &SQLLogger{SQLLevel}
}

func (l SQLLogger) Printf(format string, msg ...any) {
	file, err := OpenLogFile(l.lv)
	if err != nil {
		return
	}
	defer file.Close()

	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)
	log.SetPrefix("[" + l.lv.Name() + "] ")
	log.SetFlags(log.LstdFlags) //| log.Lmsgprefix
	log.Printf(format, msg...)
}
