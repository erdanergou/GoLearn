package main

import "fmt"

// goroutine

func hello() {
	fmt.Println("hello")
}

func main() {
	hello()
	fmt.Println("main")
}
