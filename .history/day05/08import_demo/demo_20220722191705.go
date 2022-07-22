package main  // 只有main包才能编译成可执行文件

import (
	"github.com/GoLearn/day05/07calc/calc"
	"fmt"
)

func main(){
	ret := calc.Add(1,2)
	fmt.Println(ret)
}