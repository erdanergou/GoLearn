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
func taijie(n int)(ret int){
	if n == 1{
		return 1
	}
}

func main() {

}
