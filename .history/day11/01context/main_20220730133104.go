package main

import (
	"context"
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

LOOP:
	for {
		select {
		case <-exitChan:
			break LOOP
		default:

		}
	}
}

func f2() {
	defer wg.Done()
	for {
		fmt.Printf("张三\n")
		time.Sleep(time.Microsecond * 500)
		if notify {
			break
		}
	}
}

func main() {
	wg.Add(1)
	// go f()
	exitChan = make(chan struct{}, 1)
	time.Sleep(5 * time.Second)
	notify = true
	exitChan <- struct{}{}
	// go f1()
	wg.Wait()
	// 如何通知子goroutine退出

	ctx, cancel := context.WithCancel(context.Background())
	go f2(ctx)
}
