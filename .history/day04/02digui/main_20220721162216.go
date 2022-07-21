package main

import (
	"fmt"
)

//递归

func jiecheng(n int) (ret int) {
	if n > 0 {
		return ret = n * jiecheng(n-1)
	}
	
}

func main() {
	fmt.Println(jiecheng(3))
}
