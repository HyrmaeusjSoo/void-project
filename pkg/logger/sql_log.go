package logger

import (
	"io"
	"log"
	"os"
	"time"
	"void-project/pkg"
)

type SQLLogger struct {
	lv Level
}

func NewSQLLogger() *SQLLogger {
	return &SQLLogger{SQL}
}

func (l SQLLogger) Printf(format string, msg ...any) {
	path := pkg.GetRootPath() + "/runtime/log/" + l.lv.Name() + "/"
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return
		}
	} else if err != nil {
		return
	}

	file, err := os.OpenFile(path+time.Now().Format(time.DateOnly)+".txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
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
