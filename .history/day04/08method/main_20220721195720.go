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
// 接收者表示的是调用该方法的具体l类型标量
func (d dog)wang() {
	fmt.Printf("%s 汪汪汪",d.name)
}

func main() {
	d1 := newDag("柴犬")
	d1.wang()
}
