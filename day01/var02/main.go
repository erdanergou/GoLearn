package main

import "fmt"

// 声明变量 推荐驼峰类型 studentName

// var name string
// var age int
// var isOk bool

// 批量声明变量
var (
	name string
	age  int
	isOk bool
)

func main() {
	// 同一作用域中不能重复声明同名变量
	name = "李想"
	age = 18
	isOk = true
	// Go语言中变量声明必须使用，不使用则无法编译通过
	fmt.Print(isOk)
	fmt.Println()
	fmt.Printf("name:%s", name) // %s占位符，使用name去替换
	fmt.Println(age)            // 打印完后自动添加换行符

	// 声明变量同时赋值
	var s1 string = "张三"  
	fmt.Println(s1)

	//类型推导(根据值判断该变量是什么类型)
	var s2 = "20"
	fmt.Println(s2)
	
	// 短变量声明(只能在函数中用)
	s3 := "哈哈哈"
	fmt.Println(s3)

	// 匿名变量（多重赋值或返回值想忽略某个值可以使用）

}
