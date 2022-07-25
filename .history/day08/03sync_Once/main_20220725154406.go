package main

import "sync"

// sync.Once
// Go语言中的sync包中提供了一个针对只执行一次场景的解决方案

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	once.Do(func() {
		close(ch2)
	})
}

func main() {
	a:=make(chan int, 100)
	b:=make(chan int, 100)
	wg.Add(3)
	go f1(a)
	go f
}
