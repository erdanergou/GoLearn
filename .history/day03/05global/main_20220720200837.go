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
var f1 = func(x,y int){
	
}

func main() {
	f1()
}