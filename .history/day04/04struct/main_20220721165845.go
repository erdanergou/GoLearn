package main

// 结构体
import (
	"fmt"
)

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

	// 匿名结构体
	var s  = struct{
		name string
		age int
	}
	s.name = "李四"
	s.age = 12
	fmt.Println(s)
}
