package main

import "fmt"

//闭包
//闭包是一个函数，这个函数包含了它外部作用域的一个变量

//底层原理
// 1.函数可作为返回值
// 2.函数内部查找变量的顺序，


func adder() func(int) int {
	var x = 100
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	ret := adder()
	ret2 := ret(200)

	fmt.Println(ret2)
}
