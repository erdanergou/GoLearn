package main

import "fmt"

// 接口
/*
 type 接口类型名 interface{
        方法名1( 参数列表1 ) 返回值列表1
        方法名2( 参数列表2 ) 返回值列表2
        …
    }


	任何类型的方法集中只要拥有该接口!'对应的全部方法'!签名。
	一个变量如果实现了所有接
    就表示它 "实现" 了该接口，无须在该类型上显式声明实现了哪个接口。
*/
// 接口是一种特殊类型,它规定了变量有哪些方法

// 定义一个叫类型
type speaker interface {
	speak() // 只要实现了speak方法的变量都是speaker类型,方法签名，可以有一个或多个
	// move(a int,b int)(x int, y int)
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

func da(s speaker) { // 接收一个speaker类型（接口类型）的变量
	// 接受一个参数,传进来什么就打什么
	s.speak() // 挨打
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	da(c1)
	da(d1)
	da(p1)

	var ss speaker //定义接口类型的变量
	ss = c1 
	fmt.Print(ss)
}
