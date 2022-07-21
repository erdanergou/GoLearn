package main

import "fmt"

// 函数定义
/*
func 函数名（参数）（返回值）{
	函数体
}

参数类型可以为函数，返回值也可以为函数
*/

// 闭包
/*
	函数和其外部变量的引用（

*/

//defer
/*
延迟调用，多用于处理资源释放

*/

// 内置函数
/*
panic 用于将程序退出

recover 

*/

func f1(){
	panic()
}



func f(name string) {
	fmt.Println(name)
}

func low(f func()) {
	f()
}

func main() {
	ret := bi(f, "张三")
	low(ret)
}

// 闭包
func bi(f func(string), name string) func() {
	return func() {
		f(name)
	}
}
