package main // 只有main包才能编译成可执行文件

import (
	"fmt"
	D:\GoCode\src\github.com\GoLearn\day05\07calc\calc.go
	"github.com/GoLearn/day05/07calc/calc"
)

func main() {
	ret := calc.Add(1, 2)
	fmt.Println(ret)
}
