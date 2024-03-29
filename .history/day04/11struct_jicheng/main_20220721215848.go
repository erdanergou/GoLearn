package main

import (
	"fmt"
	"go/doc"
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
	animal
}

//给狗实现叫的方法

func (d dog) wang() {
	fmt.Printf("%s 在叫叫汪汪汪", d.name)
}

func main() {
	d1 := doc{
		name:"金毛"
	}
}
