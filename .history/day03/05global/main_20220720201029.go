package main

import "fmt"

// 函数作用域

var x = 100 //定义一个全局变量

//定义一个函数
func f1() {
	// 函数中查找变量的顺序
	// 1.先从函数内部查找
	// 2.找不到找函数外，即全局，一直找到全局变量

	// 函数内部定义的变量无法再函数外使用
	fmt.Println(x)
}

// 匿名函数
// 函数内部无法定义有名函数

func main() {
	f1()
	var f2 = func(x, y int) {
		fmt.Println("123")
	}
	f2(1,2)
	// 如果是只调用一次的函数，还可以简写为立即执行函数
	func()  {
		fmt.Println("")
	}
}
