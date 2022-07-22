package main

import (
	"fmt"
	"os"
)

//打开文件

func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file err %v\n", err)
		return
	}
	//记得关闭文件
	defer file.Close()

	// 读取文件
	// var tmp = make([]byte,128) // 指定读的长度
	var tmp [128]byte
	for
}
