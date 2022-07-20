package main

import (
	"fmt"
)

//复习

func main() {
	// var ar [3]int
	// fmt.Println(ar)
	// ar = [3]int{1, 2, 3}
	// fmt.Println(ar)

	// var ar2 = [...]int{1, 3, 3, 4, 5, 6, 7}
	// fmt.Println(ar2)

	// var ar3 = [...]int{1: 3, 2: 19}
	// fmt.Println(ar3)

	// // 全局：
	// // var arr0 [5]int = [5]int{1, 2, 3}
	// // var arr1 = [5]int{1, 2, 3, 4, 5}
	// // var arr2 = [...]int{1, 2, 3, 4, 5, 6}
	// // var str = [5]string{3: "hello world", 4: "tom"}
	// // // 局部：
	// // a := [3]int{1, 2}           // 未初始化元素值为 0。
	// // b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
	// // c := [5]int{2: 100, 4: 200} // 使用索引号初始化元素。

	// // 二维数组

	// var a1 [3][3]int
	// fmt.Println(a1)
	// a1 = [3][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	// var a2 = [...][2]int{} // 只有外层可以用...
	// fmt.Println(a2)

	// // 	全局
	// //     var arr0 [5][3]int
	// //     var arr1 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	// //     局部：
	// //     a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// //     b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。
	// // }

	// //数组是值类型

	// x := [3]int{1, 2, 3}
	// y := x
	// y[1] = 200

	// fmt.Println(x)
	// f1(x)
	// fmt.Println(x)

	// 切片(slice)

	// var a []int

	// fmt.Println(a == nil) // 没有分配内存
	// a = []int{1, 2, 3}
	// fmt.Println(a)
	// a2 := make([]int, 2, 10)
	// fmt.Println(a2)
	// a3 := make([]int,0,4)
	// fmt.Println(a3 == nil)

	s1 := []int{1, 2, 3}
	s2 := s1
	var s3 []int
	s3 = make([]int, 3, 3)
	copy(s3, s1)
	fmt.Println(s2)
	s2[1] = 100
	fmt.Println(s2)
	fmt.Println(s1)
	fmt.Println(s3)

	// var s1 []int
	// s1 = append(s1, 2)  //append会自动初始化


	//指针
	// Go里面的指针只能读不能改，不能修改指针变量指向的地址

	addr := "山东"
	addp := &addr
	fmt.Println(addp)
	fmt.Printf("%T\n",addp)

}

func f1(a [3]int) {
	// go语言中函数传递的都是值
	a[0] = 199
}
