package main

import (
	"fmt"
	"sync"
)

// channel练习

var ch1 chan int
var ch2 chan int
var wg sync.WaitGroup

func f1(ch chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	
}

func f2(ch chan int) {
	defer wg.Done()
	for x := range ch {
		ch2 <- x * x
	}
}

func main() {
	ch1 = make(chan int, 10)
	ch2 = make(chan int, 20)
	wg.Add(2)
	go f1(ch1)
	go f2(ch1)
	for ret := range ch2 {
		fmt.Println(ret)
	}
	wg.Wait()
}
