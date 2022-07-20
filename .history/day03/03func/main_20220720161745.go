package main

import "fmt"

func f1(s1 string) {
	fmt.Println(s1)
}

func f2(a, b int) int {
	return a + b
}

// 可变参数，必须用在最后
func f3(a int, b ...int) {
	fmt.Printf("%T", b) // b是slice类型的切片
}



func main() {
	f1("asd")
	a := f2(3, 4)
	fmt.Println(a)
}
