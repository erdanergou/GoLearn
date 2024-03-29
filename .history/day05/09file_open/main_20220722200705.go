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
	for {
		n, err := file.Read(tmp[:])
		if err != nil {
			fmt.Printf("open file err %v\n", err)
			return
		}
		fmt.Printf("读了%d个字符\n", n)
		fmt.Println(string(tmp[:n-1]))
		if n < 128 {
			return
		}
	}
}


func readFromFile1()