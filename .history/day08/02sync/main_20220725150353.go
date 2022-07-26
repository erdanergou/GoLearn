package main

import "sync"

//锁
var x int = 0
var wg sync.WaitGroup

func add() {
	x++
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
}