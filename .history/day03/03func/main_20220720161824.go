package main

import "fmt"

func f1(s1 string) {
	fmt.Println(s1)
}

func f2(a, b int) int {
	return a + b
}

// 可变参数，必须用在最后参数上
func f3(a int, b ...int) {
	fmt.Printf("%T", b) // b是slice类型的切片
}


//返回多个
func f4(x,y int)(sub int,sum int){
	sub 
}


func main() {
	f1("asd")
	a := f2(3, 4)
	fmt.Println(a)
}
