package main // 只有main包才能编译成可执行文件

import (
	"fmt"
	"code.learn.com/Golearn/day05/"
)

func main() {
	ret := calc.Add(1, 2)
	fmt.Println(ret)
}
