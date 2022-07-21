package main

import (
	"fmt"
)

//递归
/*
要有明确的退出条件
*/
// ti'gui
func jiecheng(n int) (ret int) {
	if n > 0 {
		return n * jiecheng(n-1)
	}
	return 1
}

func main() {

}
