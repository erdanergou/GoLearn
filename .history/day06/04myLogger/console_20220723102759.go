package myLogger

import "fmt"

// 往终端写日志相关内容

// Logger日志结构体
type Logger struct {
}

 // NewLog 构造函数
func NewLog() Logger {
	return Logger{}
}


func (l Logger)Debug(msg string){
	fmt.Printf()
}