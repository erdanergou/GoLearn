package main

import "fmt"

// make()函数创造切片

// 切片的本质
// 切片就是一个框,框住了一块连续的内存。
// 切片是引用类型，真正的数据都保存在底层数组中

func main() {

	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n", s2, len(s2), cap(s2))

	// 要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。
	// var s11 []int         //len(s11)=0;cap(s11)=0;s11==nil
	// s22 := []int{}        //len(s22)=0;cap(s22)=0;s22!=nil
	// s33 := make([]int, 0) //len(s33)=0;cap(s33)=0;s33!=nil

	// 切片的赋值
	s3 := []int{1,3,5}
	s4 := s3
	fmt.Println(s3,s4)
	s3[0] = 1000
	fmt.Println(s3,s4)

	// 切片的遍历
	// 1.索引遍历
	for i:=0;i<len(s3);i++{
		fmt.Println(s3[i])
	}
	// 2.forrange遍历
	for i, v := range s3{
		fmt.Println(i, v)
	}

}
