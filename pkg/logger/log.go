package logger

import (
	"io"
	"log"
	"os"
	"time"
	"void-project/pkg"
)

func Log(lv Level, msg string) (err error) {
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

	file, err := os.OpenFile(path+time.Now().Format(time.DateOnly)+".txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)
	log.SetPrefix("[" + lv.Name() + "] ")
	log.SetFlags(log.LstdFlags) //| log.Lmsgprefix
	log.Println("-", msg)
	return
}

func LogDebug(msg string) error {
	return Log(Debug, msg)
}

func LogInfo(msg string) error {
	return Log(Info, msg)
}

func LogWarn(msg string) error {
	return Log(Warn, msg)
}

func LogError(msg string) error {
	return Log(Error, msg)
}

func LogSQL(msg string) error {
	return Log(SQL, msg)
}

// 清空日志文件
func ClearLog(lv Level) error {
	return os.RemoveAll(pkg.GetRootPath() + "/runtime/log/" + pkg.IfElse(lv != 0, lv.Name()+"/", ""))
}
