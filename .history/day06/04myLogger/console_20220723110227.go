package myLogger

import (
	"fmt"
	"strings"
	"time"
)

// 往终端写日志相关内容

type LogLevel uint16

const (
	
	DEBUG LogLevel = iota
	TARCE
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
		return DEBUG, nil
	}
}

// NewLog 构造函数
func NewLog(levelStr string) Logger {
	return Logger{}
}

func (l Logger) Debug(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [DEBUG] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Info(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Info] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Warning(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Warning] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Error(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [ERROR] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Fatal(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [FATAL] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}
