package main

import "fmt"

// 空接口（没有必要起名字，通常定义成下面的格式）
// 所有的类型都实现了空接口，也就是任意类型的变量都能保存到空接口中

//interface：关键字
//interface{} :空接口类型

func show(a interface{}){
	fmt.Printf("type %T")
}

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "张三"
	m1["age"] = 22
	m1["merried"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap", "打篮球"}
	fmt.Println(m1)

	show(false)
	show(nil)
}
