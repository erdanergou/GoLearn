package main

import "fmt"

//构造函数
type person struct {
	name string
	age  int
}

func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("张三",18)
	p2 := newPerson("李四",18)
	fmt.Println(p1,p2)
}
