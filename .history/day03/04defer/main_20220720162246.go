package main

import "fmt"

//defer

func deferDemo() {
	fmt.Println("嘿嘿嘿")
	defer fmt.Println("笑")  //defer 把ta
	fmt.Println("呵呵")
}

func main() {
	deferDemo()
}
