package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
goroutine的启动
将要并发执行的任务包装为一个函数，调用函数的时候在前面加上go关键字，就能够开启一个goroutine
goroutine对应的函数执行完，该goroutine就结束了
程序启动的时候会自动创建一个goroutine去执行main函数
main函数结束了，那么程序就结束了，由该程序启动的所有goroutine都结束了


sync.WaitGroup
var wg sync.WaitGroup

wg.Add() 计数器加一
wg.Done() 计数器减一
wg.Wait() 等


goroutine的本质
goroutine的调度模型GMP
goroutine是用户态线程，比内核态线程更轻量级，初始时只占用2kb的栈空间


channel是一种引用类型，需要通过make进行初始化（slice，map，channel都需要通过make进行初始化）
channel声明 var ch chan元素类型
channel初始化 ch = make(chan 元素类型，[缓冲区大小])

channel操作:
	发送:ch <- 100
	接受：x := <- ch
	关闭:close(ch) 非必须

							通道异常情况总结
通道	nil		非空		空的		满了		没满
接收	阻塞	接收值		阻塞		接收值		接收值
发送	阻塞	发送值		发送值		阻塞		发送值
关闭	panic	关闭成功	关闭成功	关闭成功	
				读完数据	返回零值	读完数据
				后返回零值				后返回零值
*/

// 往通道中发送值
func sendNum(ch chan<- int) {
	for {
		x := rand.Intn(10)
		ch <- x
		time.Sleep(3 * time.Second)
	}
}

func main() {
	ch := make(chan int, 10)
	go sendNum(ch)
	for {
		x, ok := <-ch      // 阻塞
		fmt.Println(x, ok) // 什么时候ok=false？当通道关闭的时候
		time.Sleep(time.Second)
	}
}
