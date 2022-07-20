package main

import "fmt"

// 函数作用域

var x = 100 //定义一个全局变量

//定义一个函数
func f1() {
	// 函数中查找变量的顺序
	// 1.先从函数内部查找
	
	fmt.Println(x)
}

func main() {
	f1()
}
