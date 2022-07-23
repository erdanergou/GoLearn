package myLogger

import (
	"fmt"
	"time"
)

// 往终端写日志相关内容

// Logger日志结构体
type Logger struct {
}

// NewLog 构造函数
func NewLog() Logger {
	return Logger{}
}

func (l Logger) Debug(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [DEBUG] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Info(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [DEBUG] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Warning(msg string) {
	now := time.Now()
	fmt.Printf("[%s] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Error(msg string) {
	now := time.Now()
	fmt.Printf("[%s] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Fatal(msg string) {
	now := time.Now()
	fmt.Printf("[%s] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}
