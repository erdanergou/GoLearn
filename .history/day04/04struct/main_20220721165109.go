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
	// 声明字段
	var zhangsan person
	zhangsan.name = "张三"
	zhangsan.age = 13
	zhangsan.gender = "male"
	zhangsan.hobby = []string{"吃", "喝"}
	fmt.Println(zhangsan)
}
