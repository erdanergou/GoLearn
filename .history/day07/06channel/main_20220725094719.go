package main

import (
	"fmt"
)

// channel
/*
go语言提倡通过通信共享内存而不是通过共享内存来实现通信
*/
var ch chan int //需要指定通道中元素的类型

func main() {
	ch = make(chan int) // 通道的初始化
	fmt.Println(ch)
}
