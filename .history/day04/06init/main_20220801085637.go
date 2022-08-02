package main

import "fmt"

//构造函数
type person struct {
	name string
	age  int
}

//构造函数(约定俗成,用new开头)
//返回的是结构体还是结构体指针
// 当结构体较大时,尽量使用结构体指针,减少内存开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}
func init() {
	fmt.Println("我是第二个初始化函数")
}
func init() {
	fmt.Println("我是第一个初始化函数")
}

func main() {
	p1 := newPerson("张三", 18)
	p2 := newPerson("李四", 18)
	fmt.Println(p1, p2)
}
