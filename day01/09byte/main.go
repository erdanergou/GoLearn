package main

import "fmt"

// byte（uint8类型）和rune（utf-8）类型

// Go语言中为了处理非Ascii码类型的字符 定义了新的rune类型

func main() {
	s := "hello 理想"
	//len()求得为byte字节的数量
	n := len(s)
	fmt.Println(n)

	// for i := 0;i<len(s);i++{
	// 	fmt.Println(s[i])
	// 	fmt.Printf("%c\n",s[i]) // %c 字符
	// }

	// for _, c := range s { // 从字符串中拿出具体的字符
	// 	// fmt.Println(c)
	// 	fmt.Printf("%c\n", c) // %c 字符
	// }

	// 字符串修改(字符串不能直接修改)
	s2 := "白萝卜"      // 白萝卜 => '白''萝''卜'
	s3 := []rune(s2) // 把字符串强制转换成了一个rune切片（分解后为字符）
	s3[0] = '红'
	fmt.Println(string(s3))

	c1 := "红"
	c2 := '红' //rune(int32)
	fmt.Printf("c1: %T c2: %T\n", c1, c2)

	// 类型转换
	num := 10
	var f float64
	f = float64(num)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
