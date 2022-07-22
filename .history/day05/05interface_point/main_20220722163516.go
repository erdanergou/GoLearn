package main

import "fmt"

// 使用值接收者和指针接收者的区别？
// 1.使用值接收者实现接口，


type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// 使用值接收者实现了接口的所有方法
// func (c cat) move() {
// 	fmt.Println("走猫步")
// }

// func (c cat) eat(food string) {
// 	fmt.Println("吃" + food)
// }

// 使用指针接收者实现接口的所有方法
func (c *cat) move() {
	fmt.Println("走猫步")
}

func (c *cat) eat(food string) {
	fmt.Println("吃" + food)
}
func main() {
	var a1 animal
	c1 := cat{"蓝短", 4}  // cat
	c2 := &cat{"英短", 4} // *cat

	// a1 = c1 // 实现animal这个接口的是cat的指针类型
	a1 = c2
	a1 = &c1
	fmt.Println(a1)

}
