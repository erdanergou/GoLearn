package main

import "fmt"

//接口保存的分为两部分：值的类型（动态类型）和值本身（动态值），这样就实现了接口变量能够存储不同的值

// 多个类型可以实现同一个接口
// 一个类型可以实现多个接口
// 

type animal interface {
	move()
	eat(string)
	suckleer
}

type suckleer interface {
	suckle(string)
}

type cat struct {
	name string
	feet int8
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡飞狗跳")
}

func (c chicken) eat(food string) {
	fmt.Println("吃" + food)
}

//cat实现animal接口
func (c cat) move() {
	fmt.Println("走猫步")
}

//cat实现animal接口
func (c cat) eat(food string) {
	fmt.Println("吃" + food)
}

//cat实现animalsuckleer接口
func (c cat) suckle(milk string) {
	fmt.Println("毛喝" + milk)
}

func main() {
	var a1 animal //定义一个接口类型变量
	c1 := cat{    // 定义一个cat类型变量
		name: "蓝短",
		feet: 4,
	}
	a1 = c1
	a1.eat("猫粮")
	fmt.Println(a1)

	c2 := chicken{
		feet: 2,
	}
	fmt.Println(c2)

}
