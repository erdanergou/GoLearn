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
	fmt.Printf("%#v\n", m1)  //值的go语言展示

	fmt.p
}
