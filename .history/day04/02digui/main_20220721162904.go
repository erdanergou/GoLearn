package main

import (
	"fmt"
)

//递归
/*
要有明确的退出条件
*/
// 递归
func jiecheng(n int) (ret int) {
	if n > 0 {
		return n * jiecheng(n-1)
	}
	return 1
}

//n阶台阶，一次可以一步或两步，有几种走法
func taijie(n int) (ret int) {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return taijie(n-1) + taijie(n-2)
	}
}

func main() {
	fmt.Println(taijie(3))
}
