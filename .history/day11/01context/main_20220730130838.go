package main

import (
	"fmt"
	"sync"
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
	var wg sync.WaitGroup
	go f()
	// 如何通知子goroutine退出
}
