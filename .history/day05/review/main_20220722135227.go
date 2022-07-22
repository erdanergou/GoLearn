package main

import (
	"encoding/json"
	"fmt"
)

//内容回顾

// 1结构体 基本数据类型
type person struct {
	name string
	age  int
	id   int64
}

type point struct {
	X int `json:"first"`
	Y int `json:"second"`
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

	p2.updateName("李晓四")

	fmt.Println(p2)

	// json序列化与反序列化

	// 序列化
	// 结构体内部的字段首字母要大写，不大写别人访问不到
	p := point{
		X: 1,
		Y: 2,
	}
	pp, err := json.Marshal(p)
	if err != nil { //如果出错了（有错误err有值）
		fmt.Printf("Marshal failed ,err %v\n", err)
	} else {
		fmt.Println(string(pp))
	}

	// 反序列化:由字符串 --> go中的结构体变量
	str1 := `"name":"张三","age":12,"id":`

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
// 结构体是值类型，当要将对象的内容进行修改时需要传递指针
func (p person) dream(str string) {
	fmt.Printf("%s 的梦想是 %s", p.name, str)
}

// 指针接收者
// 1.需要修改结构体变量的值时要使用指针接收者
// 2.结构体本身较大，拷贝的内存开销比较大时也要使用指针接收者
// 3.保持一致性：如果有一个方法使用了指针接收者其他方法也要使用以保证一致性
func (p *person) updateName(name string) {
	p.name = name
}

// 结构体嵌套
type addr struct {
	province string
	city     string
}

type student struct {
	addr // 匿名嵌套结构体
	// address addr
	name string
}
