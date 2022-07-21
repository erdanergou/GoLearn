package main

import "fmt"

//结构体是值类型

type person struct {
	name, gender string
}

func f(x person) {
	x.gender = "female"
}

func main() {
	var p person
	p.name = "张三"
	p.gender = "male"
	f(p)
	fmt.Println(p.gender)
}
