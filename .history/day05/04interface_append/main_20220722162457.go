package main

import "fmt"

//接口保存的分为两部分：值的类型（动态类型）和值本身（动态值），这样就实现了

type animal interface {
	move()
	eat(string)
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

func (c cat) move() {
	fmt.Println("走猫步")
}

func (c cat) eat(food string) {
	fmt.Println("吃" + food)
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
