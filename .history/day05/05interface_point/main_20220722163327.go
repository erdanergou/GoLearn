package main

import "fmt"

// 使用值接收者和指针接收者的区别？

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

func (c *cat) move() {
	fmt.Println("走猫步")
}

func (c cat) eat(food string) {
	fmt.Println("吃" + food)
}
func main() {
	var a1 animal
	c1 := cat{"蓝短", 4}  // cat
	c2 := &cat{"英短", 4} // *cat

	a1 = c1
	a1 = c2
	fmt.Println(a1)

}
