package logger

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
	"void-project/pkg"
)

var (
	Mode    = "" //模式，release=发布模式、dev=开发模式。发布模式只写入文件，开发模式写入文件同时控制台输出
	logFile = make(map[Level]*Logger, 10)
)

// 初始化日志
func InitLogger(mode string) {
	Mode = mode
}

// 自定义日志结构体
type Logger struct {
	valid bool
	lv    Level    // 日志级别。写入文件时会按照 runtime/log/[级别]/[日期].log切割
	file  *os.File // 写入的文件
}

// 实例化Logger
func NewLogger(lv Level) *Logger {
	l, ok := logFile[lv]
	if !ok {
		l = &Logger{false, lv, nil}
		logFile[lv] = l
	}
	return l
}

// 实例化SQLLogger
// 目前给GORM用
func NewSQLLogger() *Logger {
	return NewLogger(SQLLevel)
}

// 实现gorm.Logger接口
func (l Logger) Printf(format string, msg ...any) {
	err := l.UseOrCreate()
	if err != nil {
		return
	}
	log.Printf(format, msg...)
}

// 实例化ServerLogger
// 目前给Gin用
func NewServerLogger() io.Writer {
	l := NewLogger(ServerLevel)
	if Mode == "release" {
		return l
	}
	return io.MultiWriter(os.Stdout, l)
}

// 实现io.Writer接口 给Gin用
func (l *Logger) Write(p []byte) (n int, err error) {
	err = l.UseOrCreate()
	if err != nil {
		return
	}
	return l.file.Write(p)
}

// 使用和创建
func (l *Logger) UseOrCreate() error {
	name := time.Now().Format(time.DateOnly) + ".log"
	if !l.valid || filepath.Base(l.file.Name()) != name {
		if l.file != nil {
			l.file.Close()
		}
		file, err := openLogFile(l.lv)
		if err != nil {
			return err
		}
		l.valid = true
		l.file = file
	}

	if Mode != "release" {
		log.SetOutput(io.MultiWriter(os.Stdout, l.file))
	} else {
		log.SetOutput(l.file)
	}
	log.SetPrefix("[" + l.lv.Name() + "] ")
	log.SetFlags(log.LstdFlags) //| log.Lmsgprefix
	return nil
}

// 打开日志文件
func openLogFile(lv Level) (file *os.File, err error) {
	sep := string(os.PathSeparator)
	path := fmt.Sprintf("%v%vruntime%vlog%v%v%v", pkg.GetRootPath(), sep, sep, sep, lv.Name(), sep)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return
		}
	} else if err != nil {
		return
	}

	return os.OpenFile(path+time.Now().Format(time.DateOnly)+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
}

// 记录日志
func Log(lv Level, msg any) (err error) {
	if lv == 0 {
		return errors.New("logger:无效的日志Level")
	}
	logMsg := ""
	if err, ok := msg.(error); ok {
		logMsg = err.Error()
	} else if err, ok := msg.(string); ok {
		logMsg = err
	} else {
		return errors.New("logger:无效的日志消息体类型")
	}

	/* file, err := openLogFile(lv)
	if err != nil {
		return
	}
	defer file.Close() */

	l, ok := logFile[lv]
	if !ok {
		l = NewLogger(lv)
	}
	err = l.UseOrCreate()
	if err != nil {
		return
	}

	_, caller1, caline1, _ := runtime.Caller(2)
	_, caller2, caline2, _ := runtime.Caller(3)
	msgBuilder := strings.Builder{}
	msgBuilder.WriteString("[caller] ")
	// 判断并排除调用链中的gin.Context
	if !(strings.Contains(caller2, "github.com/gin-gonic/gin") && strings.Contains(caller2, "context.go")) {
		msgBuilder.WriteString(pkg.SubPath(caller2))
		msgBuilder.WriteRune(':')
		msgBuilder.WriteString(strconv.Itoa(caline2))
		msgBuilder.WriteString(" -> ")
	}
	msgBuilder.WriteString(pkg.SubPath(caller1))
	msgBuilder.WriteRune(':')
	msgBuilder.WriteString(strconv.Itoa(caline1))
	msgBuilder.WriteString(" [message] ")
	msgBuilder.WriteString(logMsg)
	log.Println(msgBuilder.String())
	// log.Println("-", logMsg)
	return
}

// 记录debug日志
func LogDebug(msg any) error {
	return Log(DebugLevel, msg)
}

// 记录info日志
func LogInfo(msg any) error {
	return Log(InfoLevel, msg)
}

// 记录warn日志
func LogWarn(msg any) error {
	return Log(WarnLevel, msg)
}

// 记录error日志
func LogError(msg any) error {
	return Log(ErrorLevel, msg)
}

// 记录sql日志
func LogSQL(msg any) error {
	return Log(SQLLevel, msg)
}

// 记录server日志
func LogServer(msg any) error {
	return Log(ServerLevel, msg)
}

// 清空日志文件
func ClearLog(lv Level) error {
	sep := string(os.PathSeparator)
	path := fmt.Sprintf("%v%vruntime%vlog%v", pkg.GetRootPath(), sep, sep, sep)
	if lv != 0 {
		path += lv.Name() + sep
	}
	return os.RemoveAll(path)
}
