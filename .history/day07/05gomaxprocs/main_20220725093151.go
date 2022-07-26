package main

import (
	"fmt"
	"runtime"
	"sync"
)

// GOMAXPROCS
var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

func main() {
	runtime.GOMAXPROCS(2) //不进行设置时，默认为CPU的
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}