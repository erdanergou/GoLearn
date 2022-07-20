package main

import "fmt"

//defer多用于函数结束之前释放资源（文件句柄，数据库连接，socket链接）

func deferDemo() {
	fmt.Println("嘿嘿嘿")
	// 当有多个defer时按倒叙执行
	defer fmt.Println("笑") //defer 把它后面的语句延迟到函数即将返回时执行
	defer fmt.Println("哭")
	fmt.Println("呵呵")
}

func main() {
	// deferDemo()
	// fmt.Println(f1()) // 5
	// fmt.Println(f2()) // 5
	// fmt.Println(f3()) // 5
	// fmt.Println(f4()) // 5

	
}

// go语言中函数的return不是原子操作，在底层分为两步：第一步返回值赋值，第二步RET返回
// 函数中如果存在defer，那么defer执行的时机是在第一步与第二步之间

func f1() int {
	x := 5
	defer func() {
		x++ // 修改的为x,不是返回值
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //返回值=5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ //修改的x
	}()
	return x // y = x = 5
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 返回值 = x = 5
}


func calc(index string,a,b int) int{
	ret := a + b
	fmt.Println(index,a,b,ret)
	return ret
}

