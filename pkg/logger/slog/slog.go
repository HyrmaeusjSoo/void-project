package slog

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

const (
	LevelInfo int = iota + 1
	LevelWarn
	LevelError
)

// 自定义slog
type slogger struct {
	valid bool
	slog  *slog.Logger // 标准库slog
	file  *os.File     // 写入的文件
}

var (
	path string  // 路径
	mode string  // 模式
	SLog slogger // slog
)

// 初始化slog
func InitSLog(logPath, logMode string) {
	path = logPath
	mode = logMode
}

// 写入日志
func Log(lv int, msg string, args ...any) {
	if name := time.Now().Format(time.DateOnly) + ".log"; !SLog.valid || filepath.Base(SLog.file.Name()) != name {
		if SLog.file != nil {
			SLog.file.Close()
		}

		if _, err := os.Stat(path); err != nil {
			if os.IsNotExist(err) {
				if err := os.MkdirAll(path, os.ModePerm); err != nil {
					return
				}
			} else {
				return
			}
		}
		f, err := os.OpenFile(path+name, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			return
		}

		var writer io.Writer
		if mode != "release" {
			writer = io.MultiWriter(os.Stdout, f)
		} else {
			writer = f
		}
		SLog = slogger{
			true,
			slog.New(slog.NewJSONHandler(writer, nil)),
			f,
		}
	}

	switch lv {
	case LevelInfo:
		SLog.slog.Info(msg, args...)
	case LevelWarn:
		SLog.slog.Warn(msg, args...)
	case LevelError:
		SLog.slog.Error(msg, args...)
	}
}

// 记录info日志
func Info(msg string, args ...any) {
	Log(LevelInfo, msg, args...)
}

// 记录warn日志
func Warn(msg string, args ...any) {
	Log(LevelWarn, msg, args...)
}

// 记录error日志
func Error(msg string, args ...any) {
	Log(LevelError, msg, args...)
}

// 日志文件信息
type LogInfo struct {
	Time   time.Time `json:"time"`   // 时间
	Level  string    `json:"level"`  // 级别
	Msg    string    `json:"msg"`    // 消息
	IP     string    `json:"ip"`     // 请求ip
	Method string    `json:"method"` // 请求方法
	URL    string    `json:"url"`    // 请求路径
}

// 日志文件
type LogFile struct {
	Date string    // 日期
	Logs []LogInfo // 日志列表
}

// 读取日志，时间段内的日志文件列表
//
//	s => 开始时间
//	e => 结束时间
func Read(s, e time.Time) ([]LogFile, error) {
	readLog := func(name string) (logFile LogFile, err error) {
		logFile = LogFile{}
		logPath := fmt.Sprintf("%v%v.log", path, name)
		if _, err = os.Stat(logPath); err != nil {
			return
		}
		f, err := os.OpenFile(logPath, os.O_RDONLY, 0666)
		if err != nil {
			return
		}
		defer f.Close()

		reader := bufio.NewReader(f)
		logs := []LogInfo{}
		for {
			l, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			if err != nil {
				return logFile, err
			}
			log := LogInfo{}
			err = json.Unmarshal(l, &log)
			if err != nil {
				break
			}
			logs = append(logs, log)
		}

		logFile.Date = name
		logFile.Logs = logs
		return
	}

	begin := time.Date(s.Year(), s.Month(), s.Day(), 0, 0, 0, 0, time.Local)
	end := time.Date(e.Year(), e.Month(), e.Day(), 23, 59, 59, 0, time.Local)
	files := []LogFile{}
	for i := begin; i.Before(end); i = i.Add(24 * time.Hour) {
		date := i.Format(time.DateOnly)
		file, err := readLog(date)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}
