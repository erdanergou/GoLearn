package main

import "fmt"

// 闭包

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}
// 定义一个函数对f2进行包装
func f3(int,int) func(){
	tmp := func()  {
		
	}
	return tmp
}
func main() {

}
