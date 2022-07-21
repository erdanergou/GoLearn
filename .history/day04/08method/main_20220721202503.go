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

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// 方法是作用域特定类型的函数
// 接收者表示的是调用该方法的具体l类型变量,多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s 汪汪汪", d.name)
}

//使用值接收者:传拷贝进去
func (p person) guonian() {
	p.age++
}

//使用指针接收者:传内存地址进去
/*什么时候使用指针接收者
1.需要修改接收者中的值
2.接收者是拷贝代价过大的大对象
3.保证一致性,如果某个方法使用了指针接收者,那么其他的方法也应该使用指针接收者

*/
func (p *person) guozhennian() {
	p.age++
}

func 

func main() {
	d1 := newDag("柴犬")
	d1.wang()

	var p = newPerson("张三", 19)
	p.guonian()
	fmt.Println(p.age)
	p.guozhennian()
}
