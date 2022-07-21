package main

import (
	"fmt"
)

// 自定义类型和类型别名

//type后面跟的是类型
type myInt int

type yourInt = int 

func main() {
	var n myInt
	n = 100
	fmt.Printf("%T\n", n)
}
