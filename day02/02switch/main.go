package main

import "fmt"

func main() {
	var n = 3
	//switch
	switch n {
	case 1:
		fmt.Println("大拇指")

	case 2:
		fmt.Println("食指")

	case 3:
		fmt.Println("中指")
	default:
		fmt.Println("无效的数字")
	}
}
