package main

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
wg.Wait() 

*/


func main() {

}
