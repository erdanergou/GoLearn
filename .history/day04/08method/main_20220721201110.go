package main

import "fmt"

//方法
// go语言中如果标识符首字母是大写的,就表示对外部可见(暴露的,公有的)

type dog struct {
	name string
}

//构造函数
func newDag(name string) dog {
	return dog{
		name: name,
	}
}

type person struct {
	name string
	age  int
}

// 方法是作用域特定类型的函数
// 接收者表示的是调用该方法的具体l类型变量,多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s 汪汪汪", d.name)
}

func (p person) guonian() {
	p.age++
}

func main() {
	d1 := newDag("柴犬")
	d1.wang()
}
