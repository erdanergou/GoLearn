package main

//锁
var x int = 0

func add() {
	x++
}

func main() {
	go add()
	go add()
}
