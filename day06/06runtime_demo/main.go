package main

import (
	"fmt"
	"path"
	"runtime"
)

// runtime
// 用Caller能够拿到谁调用了这一行代码，包括文件名，函数名还有行号
func getInfo(n int) {
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Printf("runtime.Caller failed \n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name() //函数名
	fmt.Println(funcName)
	fmt.Println(file)
	fmt.Println(path.Base(file))
	fmt.Println(line)
}

func main() {
	getInfo(1)
}
