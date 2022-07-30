package main

import (
	"fmt"
	"sync"
	"time"
)

//context
var wg sync.WaitGroup

func f() {
	defer wg.Done()
	for {
		fmt.Printf("张三")
		time.Sleep(time.Microsecond * 500)
	}
}

func main() {
	wg.Add(1)
	go f()
	wg.Wait()
	// 如何通知子goroutine退出
}
