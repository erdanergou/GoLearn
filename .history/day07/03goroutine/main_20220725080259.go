package main

import "fmt"

// goroutine

func hello() {
	fmt.Println("hello")
}

func main() {
	go hello()  // 开启一个单纯的
	fmt.Println("main")
}
