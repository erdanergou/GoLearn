package main

import "fmt"

func main() {
	var (
		a = 5
		b = 2
	)
	// 算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a++ //a = a + 1
	b-- // b = b - 1

	// 关系运算符
	fmt.Println(a == b) // go语言是强类型，相同类型的变量才能比较
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a < b)

	// 逻辑运算符 || && not
	// 如果年龄大于18岁且年龄小于60
	age := 28
	if age > 18 && age < 60 {
		fmt.Println("年轻打工人")
	} else {
		fmt.Println("不干活")
	}

	// 位运算符 针对二进制数
	// 5  101
	// 2  010
	// 按位与 两位均为1才为1
	fmt.Println(5 & 2) // 000
	// 按位或 两位有一个为1就为1
	fmt.Println(5 | 2) //111
	// 按位异或 两位不一样则为1
	fmt.Println(5 ^ 2) //111
	// << 二进制数左移
	fmt.Println(5 << 1)
	// >> 二进制右移
	fmt.Println(5 >> 1)
}
