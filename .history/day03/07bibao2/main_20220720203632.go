package main

import "fmt"

//闭包
//闭包是一个函数，这个函数包含了它外部作用域的一个变量



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
