package main

import "fmt"

//结构体是值类型

type person struct {
	name, gender string
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
	
	var p2 = new(person) //得到的是结构体的地址
	fmt.Printf("%T\n", p2)


	var a int
	a = 100
	b := &a
	fmt.Printf("type a:%T type b:%T",a,b)
	fmt.Printf("%p\n",&a)
	fmt.Printf("%p\n",b)
}
