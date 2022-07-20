package main

import "fmt"

func f1(s1 string) {
	fmt.Println(s1)
}

func f2(a, b int) int {
	return a + b
}

// 可变参数，必须用在最后参数上
func f3(a int, b ...int) {
	fmt.Printf("%T", b) // b是slice类型的切片
}

//返回多个
func f4(x, y int) (sub int, sum int) {
	sub = x - y
	sum = x + y
	return
}

func main() {
	// f1("asd")
	// a := f2(3, 4)
	// fmt.Println(a)
	// Go中不允许在函数内再声明函数
	// func f8()  {

	// }

	a := f5
	b := f6
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}

// 函数类型
func f5() {
	fmt.Println("Hello")
}

func f6() int {
	return 10
}

// 函数也可以作为参数类型
func f7(x func())  {
	
}