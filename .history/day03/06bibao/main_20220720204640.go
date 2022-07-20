package main

import "fmt"

// 闭包

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 定义一个函数对f2进行包装
func f3(f func(int, int), m, n int) func() {
	tmp := func() {
		f(m, n)
	}
	return tmp
}

func f4(f func(int, int), x, y int) func() {
	tmp := func() {
		// fmt.Println(x + y)
		// f2
		f(x, y)
	}
	return tmp
}

func main() {
	ret := f4(f2, 100, 200) // 把原来
	f1(ret)
}
