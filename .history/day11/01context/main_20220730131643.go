package main

import (
	"fmt"
	"sync"
	"time"
)

//context
var wg sync.WaitGroup

var notify bool

var exitChan chan struct{}

func f() {
	defer wg.Done()
	for {
		fmt.Printf("张三\n")
		time.Sleep(time.Microsecond * 500)
		if notify {
			break
		}
	}
}

func f1() {
	defer wg.Done()

	for {
		select {
		case v := <-exitChan:
			fmt.Println(v)
			break
		default:
			fmt.Println()
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(5 * time.Second)
	notify = true
	wg.Wait()
	// 如何通知子goroutine退出
}
