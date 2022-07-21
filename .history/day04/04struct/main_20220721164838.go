package main

// 结构体

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var zhangsan person
	zhangsan.name = "张三"
	zhangsan.age = 13
	zhangsan.gender = "male"
	zhangsan.hobby = {}
}
