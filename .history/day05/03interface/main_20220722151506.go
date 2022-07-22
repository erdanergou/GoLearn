package main

import "fmt"

// 接口
// 接口是一种特殊类型,它规定了变量有哪些方法


// 定义一个叫类型
type speaker interface{
	
}

type cat struct{}

type dog struct{}

type person struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵~~~")
}

func (d dog) speak() {
	fmt.Println("汪汪汪~~~")
}

func (p person) speak() {
	fmt.Println("啊~")
}

func da(x) {
	// 接受一个参数,传进来什么就打什么
	x.speak() // 挨打
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person
}
