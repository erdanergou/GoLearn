package main

import (
	"fmt"
	"sync"
)

//锁
var x int = 0
var wg sync.WaitGroup
var lock sync.Mutex // 互斥锁

func add() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		x++
	}
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
