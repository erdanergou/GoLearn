package main

import "fmt"

//方法

type dog struct {
	name string
}

//构造函数
func newDag(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用域特定类型的函数 
func (d dog)wang() {
	fmt.Printf("%s 汪汪汪",d.name)
}

func main() {
	d
}
