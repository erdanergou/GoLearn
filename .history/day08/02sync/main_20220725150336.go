package main

//锁
var x int = 0
var wg 

func add() {
	x++
}

func main() {
	go add()
	go add()
}
