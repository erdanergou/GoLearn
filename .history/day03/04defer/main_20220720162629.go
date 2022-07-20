package main

import "fmt"

//defer多用于函数结束之前释放资源

func deferDemo() {
	fmt.Println("嘿嘿嘿")
	// 当有多个defer时按倒叙执行
	defer fmt.Println("笑") //defer 把它后面的语句延迟到函数即将返回时执行
	defer fmt.Println("哭")
	fmt.Println("呵呵")
}

func main() {
	deferDemo()
}
