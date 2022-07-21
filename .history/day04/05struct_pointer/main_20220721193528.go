package main

import "fmt"

//结构体是值类型

type person struct {
	name, gender string
}

// 结构体占用一块连续的内存空间
type x struct{
	a int8
	b int8
	c int8
}

// go语言中函数参数永远是拷贝
func f(x person) {
	x.gender = "female" //修改的是副本的gender，不会对原始数据进行修改
}

func f2(x *person) {
	// *&(x).gender = "female"   // 根据内存地址找到原变量，修改的就是传入的值
	x.gender = "female" // 语法糖,自动根据指针找对应变量
}



func main() {
	var p person
	p.name = "张三"
	p.gender = "male"
	f(p)
	fmt.Println(p.gender)

	f2(&p)
	fmt.Println(p.gender)

	// 结构体指针1
	var p2 = new(person) //得到的是结构体的地址
	p2.gender = "female"
	*&(p2).name = "李四"
	fmt.Printf("%T\n", p2)

	// 2.1 key-value初始化
	var p3 = person{
		name:   "王五",
		gender: "male",
	}
	fmt.Printf("%#v", p3)

	// 2.2 使用值列表的形式初始化,(若不加标识,之的顺序要与结构体定义时字段的顺序一致)
	p4 := person{
		gender: "female",
		name:   "赵六",
	}
	fmt.Printf("%#v", p4)


	// 证明结构体占用连续的内存
	m := x{
		a : int8(10),
		b : int8(20),
		c : int8(30),
	}
	fmt.Printf()


}
