package main

import "sync"

//锁
var x int = 0
var wg sync.WaitGroup

func add() {
	for i:=0;i<10;
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
}
