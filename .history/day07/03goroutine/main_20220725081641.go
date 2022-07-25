package main

import (
	"fmt"
	"time"
)

// goroutine
/*
goroutine什么时候结束
goroutine对应的函数结束了，goroutine结束了
main函数所在的

*/



func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动之后会创建一个主goroutine去执行
func main() {
	for i := 0; i < 10; i++ {
		go hello(i) // 开启一个单独的goroutine去执行hello函数（任务）
		// 匿名函数
		go func(i int) {
			fmt.Println(i) // 如果不传入参数，打印的i为for循环的i，可能会出现打印重复现象
		}(i)
	}
	fmt.Println("main")
	// time.Sleep(time.Second * 1)
	// main函数结束了 由main函数启动的goroutine也都结束了

}
