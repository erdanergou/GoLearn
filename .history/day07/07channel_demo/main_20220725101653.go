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
	for{
		
	}

}

func f2() {
	defer wg.Done()
	for x := range ch1 {
		ch2 <- x * x
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
