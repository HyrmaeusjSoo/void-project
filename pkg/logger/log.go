package logger

import (
	"errors"
	"io"
	"log"
	"os"
	"time"
	"void-project/pkg"
)

func OpenLogFile(lv Level) (file *os.File, err error) {
	path := pkg.GetRootPath() + "/runtime/log/" + lv.Name() + "/"
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return
		}
	} else if err != nil {
		return
	}

	return os.OpenFile(path+time.Now().Format(time.DateOnly)+".txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
}

func Log(lv Level, msg any) (err error) {
	logMsg := ""
	if err, ok := msg.(error); ok {
		logMsg = err.Error()
	} else if err, ok := msg.(string); ok {
		logMsg = err
	} else {
		return errors.New("无效的日志消息体类型")
	}

	file, err := OpenLogFile(lv)
	if err != nil {
		return
	}
	defer file.Close()

	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)
	log.SetPrefix("[" + lv.Name() + "] ")
	log.SetFlags(log.LstdFlags) //| log.Lmsgprefix
	log.Println("-", logMsg)
	return
}

func LogDebug(msg any) error {
	return Log(DebugLevel, msg)
}

func LogInfo(msg any) error {
	return Log(InfoLevel, msg)
}

func LogWarn(msg any) error {
	return Log(WarnLevel, msg)
}

func LogError(msg any) error {
	return Log(ErrorLevel, msg)
}

func LogSQL(msg any) error {
	return Log(SQLLevel, msg)
}

// 清空日志文件
func ClearLog(lv Level) error {
	return os.RemoveAll(pkg.GetRootPath() + "/runtime/log/" + pkg.IfElse(lv != 0, lv.Name()+"/", ""))
}
