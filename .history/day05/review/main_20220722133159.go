package main

import "fmt"

//内容回顾

// 1结构体 基本数据类型
type person struct {
	name string
	age  int
	id   int64
}

func main() {
	// 2.匿名结构体 多用于临时场景，
	var s = struct {
		x int
		y int
	}{10, 20}
	fmt.Println(s)

	p1 := person{ //结构体的实例化
		name: "张三",
		age:  15,
	}
	fmt.Println(p1)

	p2 := newPerson("李四", 20, 2)
	fmt.Println(p2)

	p2.dream("躺着把钱赚了")
}

// 结构体的构造函数,返回值是对应的结构体类型
func newPerson(name string, age int, id int) person {
	return person{
		name: name,
		age:  age,
		id:   int64(id),
	}
}

// 方法是有接收者的函数，接收者指的是哪个类型的变量可以调用这个函数
// 指定了接收者之后，只有接收者这个类型可以调用
// 结构体是值类型，当要
func (p person) dream(str string) {
	fmt.Printf("%s 的梦想是 %s", p.name, str)
}
