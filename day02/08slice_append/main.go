package main

import "fmt"

// append() 为切片追加元素

func main() {
	s1 := []string{"泰安", "济南", "青岛"}
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n",s1,len(s1),cap(s1))
	// s1[3] = "烟台"  //会导致编译错误，索引越界
	// fmt.Println(s1)
	
	// 调用append函数必须用原来的切片变量接收返回值
	s1 = append(s1, "烟台")  // append追加元素，原来的底层数组放不下时，
							// Go底层就会把底层数组进行更换，必须用变量接收append的返回值
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n",s1,len(s1),cap(s1))
	
	s2 := []string{"日照","临沂","济宁","德州"}
	s1 = append(s1, s2...) // ...表示拆开，
	
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n",s1,len(s1),cap(s1))


	//关于append删除切片中的某个元素
	aa1 := [...]int{1,3,5,7,9,11,13,15,17,19}
	ss1 := aa1[:]

	//删除3
	ss1 = append(ss1[:1],ss1[2:]...)
	fmt.Println(ss1)
	fmt.Println(aa1)
}
