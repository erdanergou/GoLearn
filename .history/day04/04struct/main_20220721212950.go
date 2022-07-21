package main

// 结构体
import (
	"fmt"
)

// 结构体遇到的问题

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// 声明一个person类型的变量p
	var zhangsan person
	// 通过字段赋值
	zhangsan.name = "张三"
	zhangsan.age = 13
	zhangsan.gender = "male"
	zhangsan.hobby = []string{"吃", "喝"}
	fmt.Println(zhangsan)
	fmt.Printf("%T\n", zhangsan)

	// 匿名结构体:多用于临时场景
	var s struct {
		name string
		age  int
	}
	s.name = "李四"
	s.age = 12
	fmt.Println(s)
	fmt.Printf("%T\n", s)
}
