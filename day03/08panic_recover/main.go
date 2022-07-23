package main

import "fmt"

//panic 和 recover
// recover()一定要搭配defer使用
// defer 一定要在可能引发panic的语句之前定义


func funcA() {
	fmt.Println("a")
}
func funcB() {
	// 刚刚打开数据库
	defer func() {
		err:=recover()
		fmt.Println(err)
		fmt.Println("释放数据库连接")
	}()
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
