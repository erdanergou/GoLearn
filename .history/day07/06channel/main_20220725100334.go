package main

import (
	"fmt"
	"sync"
)

// channel
/*
go语言提倡通过通信共享内存而不是通过共享内存来实现通信
通道必须使用make函数初始化才能使用

通道中放的尽量小，如果大就放地址		

通道的操作：
	1.发送 ch <- 19
	2.接受 x := <- ch
	3.关闭 close
*/
var ch chan int //声明通道类型的变量，需要指定通道中元素的类型
var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(ch)
	ch = make(chan int) // 不带缓冲区的通道的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-ch
		fmt.Printf("后台goroutine接收到通道ch中的数据%d了\n", x)
	}()
	// 发送,将一个值发送到通道中
	ch <- 10 // 把10发送到ch中
	fmt.Println("数据发送到通道ch中了。。。")
	fmt.Println(ch)
	wg.Wait()
}

func BufChannel() {
	fmt.Println(ch)
	ch = make(chan int, 16) // 不带缓冲区的通道的初始化
	// 发送,将一个值发送到通道中
	ch <- 10 // 把10发送到ch中
	fmt.Println("数据发送到通道ch中了。。。")
	x := <- ch
	pr
}

func main() {
	BufChannel()
}
