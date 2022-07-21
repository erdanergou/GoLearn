package main

import "fmt"

//结构体是值类型

type person struct {
	name, gender string
}

func f(x person) {
	x.gender = "female"  //修改的是副本的gender，不会对原始
}

func main() {
	var p person
	p.name = "张三"
	p.gender = "male"
	f(p)
	fmt.Println(p.gender)
}
