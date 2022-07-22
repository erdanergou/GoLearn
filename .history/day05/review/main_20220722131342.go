package main

import "fmt"

//内容回顾

// 1结构体 基本数据类型
type person struct {
	name string
	age  int
	id   int64
}

// 2.匿名结构体 多用于临时场景，

func main() {
	var s = struct {
		x int
		y int
	}{10, 20}
	fmt.Println(s)
}
