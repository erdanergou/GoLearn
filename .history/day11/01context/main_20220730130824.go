package main

import (
	"fmt"
	"time"
)

//context

func f() {
	for {
		fmt.Printf("张三")
		time.Sleep(time.Microsecond * 500)
	}
}

func main() {
	var wg 
	go f()
	// 如何通知子goroutine退出
}
