package main

import (
	"fmt"
)

//结构体模拟实现其他语言中的继承

type animal struct {
	name string
}

//给动物一个移动的方法
func (a animal) move() {
	fmt.Printf("%s 会动", a.name)
}

type dog struct {
	feet uint8
	animal  //animal拥有的方法dog此时  
}

//给狗实现叫的方法

func (d dog) wang() {
	fmt.Printf("%s 在叫叫汪汪汪", d.name)
}

func main() {
	d1 := dog{
		animal: animal{
			name: "哈士奇",
		},
		feet: 4,
	}
	fmt.Println(d1)
	d1.move()
	d1.wang()
}
