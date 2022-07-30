package main

import (
	"fmt"
	"sync"
	"time"
)

//context
var wg sync.WaitGroup

func f() {
	for {
		fmt.Printf("张三")
		time.Sleep(time.Microsecond * 500)
	}
	defer wg.Done()
}

func main() {

	go f()
	// 如何通知子goroutine退出
}
