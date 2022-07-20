package main

import "fmt"

func main() {
	fmt.Print("山东")
	fmt.Print("泰安")
	fmt.Println("~~~~~~~~~~~~~~")
	fmt.Println("山东")
	fmt.Println("泰安")
	// Printf("标准化字符串"，值)

	var m1 = make(map[string]int, 1)
	m1["张三"] = 19
	fmt.Printf("%v\n", m1)
	fmt.Printf("%#v\n", m1) //值的go语言展示

	fmt.Printf("%d%%\n", 90) //90%
	// 整数->字符
	fmt.Printf("%q\n", 80)
	// 浮点数和复数
	fmt.Printf("%g\n", 3.141692654)
	// 字符串
	fmt.Printf("%q\n", "张三有三")

	fmt.Printf("%2.2f\n", 12.245)

	//获取用户输入
	var s string
	fmt.Scan(&s)
	fmt.Println(s)

	var (
		name  string
		age   int
		class string
	)
	fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Println(name, age, class)
	fmt.Scanf("%s %d %s\n", &name, &age, &class)
}
