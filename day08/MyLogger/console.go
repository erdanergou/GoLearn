package myLogger

import (
	"fmt"
	"time"
)

// 往终端写日志相关内容

// NewLog 构造函数
func NewConsoleLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(level LogLevel) bool {
	return c.Level <= level
}

func (c ConsoleLogger) log(lv LogLevel, msg string, arg ...interface{}) {
	if c.enable(lv) {
		fotmat := fmt.Sprintf(msg, arg...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s: %s: %d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, fotmat)
	}
}

func (c ConsoleLogger) Debug(msg string, arg ...interface{}) {

	c.log(DEBUG, msg, arg...)
}

func (c ConsoleLogger) Info(msg string, arg ...interface{}) {
	c.log(INFO, msg, arg...)
}

func (c ConsoleLogger) Warning(msg string, arg ...interface{}) {
	c.log(WARNING, msg, arg...)
}

func (c ConsoleLogger) Error(msg string, arg ...interface{}) {
	c.log(ERROR, msg, arg...)
}

func (c ConsoleLogger) Fatal(msg string, arg ...interface{}) {
	c.log(FATAL, msg, arg...)
}
