package main

import (
	"fmt"
	"sync"
)

// channel练习

var ch1 chan int
var ch2 chan int
var wg sync.WaitGroup

func f1() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch1 <- i
	}

}

func f2() {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x
	}
}

func main() {
	ch1 = make(chan int, 10)
	ch2 = make(chan int, 20)
	wg.Add(2)
	go f1()
	go f2()
	for ret := range ch2 {
		fmt.Println(ret)
	}
	wg.Wait()
}
