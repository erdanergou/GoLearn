package main

import "fmt"

//构造函数
type person struct {
	name string
	age  int
}

//构造函数
//返回的是结构体还是结构体指针
// 当结构体较大时,尽量使用结构体指针,减少内存
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("张三", 18)
	p2 := newPerson("李四", 18)
	fmt.Println(p1, p2)
}
