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
// 方法使用zhi'jie'shou'z
func (c cat) move() {
	fmt.Println("走猫步")
}

func (c cat) eat(food string) {
	fmt.Println("吃" + food)
}
func main() {

}
