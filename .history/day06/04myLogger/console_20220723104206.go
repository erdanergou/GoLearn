package myLogger

import (
	"fmt"
	"os"
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
	
	fmt.Printf("[%s] %s",n)
}

func (l Logger) Info(msg string) {
	fmt.Fprintf(os.Stdout, msg)
}

func (l Logger) Warning(msg string) {
	fmt.Fprintf(os.Stdout, msg)
}

func (l Logger) Error(msg string) {
	fmt.Fprintf(os.Stdout, msg)
}

func (l Logger) Fatal(msg string) {
	fmt.Fprintf(os.Stdout, msg)
}
