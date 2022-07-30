package main

import (
	"fmt"
	"sync"
	"time"
)

//context
var wg sync.WaitGroup

var notify bool

func f() {
	defer wg.Done()
	for {
		fmt.Printf("张三")
		time.Sleep(time.Microsecond * 500)
		if notify {
			break
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	notify = true
	wg.Wait()
	// 如何通知子goroutine退出
}
