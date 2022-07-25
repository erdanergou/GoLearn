package main

import (
	"fmt"
	"sync"
)

//锁
var x int = 0
var wg sync.WaitGroup

func add() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
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
