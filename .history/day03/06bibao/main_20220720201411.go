package main

import "fmt"

// 闭包

func f1(f func()) {
	
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}
func main() {

}
