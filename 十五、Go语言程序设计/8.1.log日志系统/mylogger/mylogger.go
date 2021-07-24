package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "panic":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

type Logger struct {
	Level LogLevel
}

func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func (l Logger) enable(logLevel LogLevel) bool {
	return l.Level <= logLevel
}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"

	}
	return "DEBUG"
}

func log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	now.Format("2006-01-02 15:04:05")
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now, getLogString(lv), funcName, fileName, lineNo, msg)
}

func (l *Logger) Debug(format string, a ...interface{}) {
	if l.enable(DEBUG) {
		log(l.Level, format, a)
	}
}

func (l *Logger) Trace(msg string) {
	if l.enable(TRACE) {
		now := time.Now()
		now.Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(2)
		fmt.Printf("[%s] [Debug] [%s:%s:%d] %s\n", now, funcName, fileName, lineNo, msg)
	}
}

func (l *Logger) Info(format string, a ...interface{}) {
	if l.enable(INFO) {
		log(l.Level, format, a)
	}
}

func (l *Logger) Warning(format string, a ...interface{}) {
	if l.enable(WARNING) {
		log(l.Level, format, a)
	}
}

func (l *Logger) Error(format string, a ...interface{}) {
	if l.enable(ERROR) {
		log(l.Level, format, a)
	}
}

func (l *Logger) Panic(format string, a ...interface{}) {
	if l.enable(FATAL) {
		log(l.Level, format, a)
	}
}

func getInfo(n int) (funcName, fileName string, lineNo int) {
	// pc表示函数的信息
	// file表示调用runtime.Caller的文件名
	// line表示行号
	// ok表示 如果能取到信息，ok就为true
	pc, file, lineNo, ok := runtime.Caller(n)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	// n传入0，表示直接调用runtime.Caller的函数
	// 假设直接调用的函数为f1, 调用f1的函数为f2, n传入1，表示调用 直接调用runtime.Caller的函数 的函数，即f2
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)

	return
}
