package main

import "fmt"

// 整形

func main(){
	var i1 = 101  // 十进制
	fmt.Printf("%d\n", i1) // 输出为十进制
	fmt.Printf("%b\n", i1) // 输出为二进制
	fmt.Printf("%o\n", i1) // 输出为八进制
	fmt.Printf("%x\n", i1) // 输出为十六进制

	var i2 = 077  // 八进制
	fmt.Printf("%d\n", i2)

	i3 := 0xff // 十六进制
	fmt.Printf("%d\n",i3)

	//查看变量类型
	fmt.Printf("%T\n",i3)

	// 声明int8类型的变量
	i4 := int8(9) // 明确指定为int8类型，否则默认为int类型
	fmt.Printf("%T\n",i4)
}