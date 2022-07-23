package myLogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

const (
	UNKNOWN LogLevel = iota
	TARCE
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

type LogLevel uint16

// Logger日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

func parseLogLevel(s string) (LogLevel, error) {
	switch strings.ToLower(s) {
	case "debug":
		return DEBUG, nil
	case "tarce":
		return TARCE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("日志级别无效")
		return UNKNOWN, err
	}
}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TARCE:
		return "TARCE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller failed \n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name() //函数名
	fileName = path.Base(file)
	return
}
