package main

import "fmt"

//内容回顾

// 1结构体 基本数据类型
type person struct {
	name string
	age  int
	id   int64
}

func main() {
	// 2.匿名结构体 多用于临时场景，
	var s = struct {
		x int
		y int
	}{10, 20}
	fmt.Println(s)

	p1 := person{ //结构体的实例化
		name: "张三",
		age:  15,
	}
	fmt.Println(p1)
}


// jie'gou't
