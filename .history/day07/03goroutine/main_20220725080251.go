package main

import "fmt"

// goroutine

func hello() {
	fmt.Println("hello")
}

func main() {
	go hello()
	fmt.Println("main")
}
