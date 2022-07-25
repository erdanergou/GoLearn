package main

import "fmt"

// goroutine

func hello() {
	fmt.Println("hello")
}

// 程序启动
func main() {
	go hello() // 开启一个单独的goroutine去执行hello函数（任务）
	fmt.Println("main")
}
