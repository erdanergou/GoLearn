package main

import (
	"fmt"
	"os"
)

// os.Args 获取命令行参数

func main() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Printf("%T\n", os.Args)
}
