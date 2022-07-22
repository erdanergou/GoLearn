package main // 只有main包才能编译成可执行文件

import (
	"fmt"

	calc "code.learn.com/Golearn/day05/07calc"
)

// 包的导入注意事项
/*
import导入语句通常放在文件开头包声明语句的下面。
导入的包名需要使用双引号包裹起来。
●包名是从$GOPATH/src/后开始计算的， 使用进行路径分隔。
●Go语言中禁止循环导入包。


*/

func main() {
	ret := calc.Add(1, 2)
	fmt.Println(ret)
}
