package main

import "fmt"

//给自定义类型加方法
// 不能给别的包里面的类型添加类型
type myInt int

func (i myInt) sayHello() {
	fmt.Println("我是一个int")
}

func main() {
	m := myInt(100)
	m.sayHello()
}
