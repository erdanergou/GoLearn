package main

import "fmt"

// 切片slice

func main() {
	// 切片的定义
	var s1 []int    //定义一个存放int类型元素的切片
	var s2 []string //定义一个存放int类型元素的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // true nil 相当于空
	fmt.Println(s2 == nil)

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"北京", "上海", "武汉"}

	// 长度和容量
	fmt.Printf("len(s1): %d cap(s1):%d\n",len(s1),cap(s1))
	fmt.Printf("len(s2): %d cap(s2):%d\n",len(s2),cap(s2))

	//2.由数组得到切片
	a1 := [...]int{1,3,5,7,9,11,13}
	s3 := a1[0:4] // 基于一个数组切割，左闭右开
	fmt.Println(a1 , s3)
	s4 := a1[:4]
	fmt.Println(s4)
	s5 := a1[3:]
	s6 := a1[:]
	fmt.Println(s5,s6)
	// 切片的容量是指底层数组的切片的第一个元素到最后的元素数量
	fmt.Printf("len(s4):%d cap(s4):%d\n",len(s4),cap(s4))
	
	// 底层数组从切片的第一个元素到最后的元素数量
	fmt.Printf("len(s5):%d cap(s5):%d\n",len(s5),cap(s5))

	//切片在切片
	s7 := s5[3:]
	fmt.Printf("len(s7):%d cap(s7):%d\n",len(s7),cap(s7))


	//
}
