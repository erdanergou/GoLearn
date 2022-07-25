package main

import (
	"fmt"
	"sync"
)

// sync.Once
// Go语言中的sync包中提供了一个针对只执行一次场景的解决方案
//注意：如果要执行的函数f需要传递参数就需要搭配闭包来使用。

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
		close(ch2) // 确保某个操作只执行一次
	})
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}
}
