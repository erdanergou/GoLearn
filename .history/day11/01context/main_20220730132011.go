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
		case  <-exitChan:
			return
		default:
			fmt.Println("张三")
		}
	}
}

func main() {
	wg.Add(1)
	// go f()
	exitChan = make(chan struct{}, 1)
	time.Sleep(5 * time.Second)
	notify = true
	x := 
	exitChan <- true
	go f1()
	wg.Wait()
	// 如何通知子goroutine退出

}
