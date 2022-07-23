package main

import "fmt"

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

func (c chicken) eat() {
	fmt.Println("吃鸡饲料")
}

func (c cat) move() {
	fmt.Println("走猫步")
}

func (c cat) eat() {
	fmt.Println("吃猫粮")
}

func main() {
	var a1 animal
	c1 := cat{
		name: "",
	}
}