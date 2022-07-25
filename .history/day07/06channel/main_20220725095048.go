package main

import (
	"fmt"
)

// channel
/*
go语言提倡通过通信共享内存而不是通过共享内存来实现通信
通道必须使用make函数初始化才能使用
*/
var ch chan int //声明通道类型的变量，需要指定通道中元素的类型

func main() {
	fmt.Println(ch)
	ch = make(chan int) // 不带缓冲区的通道的初始化
	// ch = make(chan int,16) // 带缓冲区的通道的初始化
	fmt.Println(ch)

	// 发送
	
}
