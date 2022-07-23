package myLogger

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// 往终端写日志相关内容

type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	TARCE
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

// Logger日志结构体
type Logger struct {
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

// NewLog 构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func (l Logger) enable(level LogLevel) bool {
	return l.Level > level
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		now := time.Now()
		fmt.Printf("[%s] [DEBUG] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		fmt.Printf("[%s] [Info] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
	now := time.Now()
	fmt.Printf("[%s] [Warning] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
	now := time.Now()
	fmt.Printf("[%s] [ERROR] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Fatal(msg string) {
	if l.enable(DEBfaUG) {
	now := time.Now()
	fmt.Printf("[%s] [FATAL] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}
