package main

import (
	"fmt"
	"time"
)

// goroutine

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动之后会创建一个主goroutine去执行
func main() {
	for i:= 0;i < 10;i++{
		go hello() // 开启一个单独的goroutine去执行hello函数（任务）

	}
	fmt.Println("main")
	time.Sleep(time.Second * 1)
	// main函数结束了 由main函数启动的goroutine也都结束了
}
