package main // 只有main包才能编译成可执行文件

import (
	"fmt"

	calc "code.learn.com/Golearn/day05/07calc"
)

// 包的导入

func main() {
	ret := calc.Add(1, 2)
	fmt.Println(ret)
}
