package main

import (
	"fmt"
)

//结构体与json
// 1. 序列化：把go语言中的结构体变量 --> json格式的字符串
// 2. 反序列化：json格式的字符串 --> go语言中能够识别的结构体变量

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{
		name: "张三",
		age:  17,
	}
	fmt.Println(p1)
}
