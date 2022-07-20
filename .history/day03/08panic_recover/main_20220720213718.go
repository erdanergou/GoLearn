package main

import "fmt"

//panic 和 recover

func funcA() {
	fmt.Println("a")
}
func funcB() {
	// 刚刚打开数据库
	
	panic("出现严重错误") //程序崩溃退出
	fmt.Println("b")
}
func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
