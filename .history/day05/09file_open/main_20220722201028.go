package main

import (
	"bufio"
	"fmt"
	"os"
)

//打开文件

func main() {
	// readFromFile1("./main.go")
}

func readFromFile1(path string) {
	file, err := os.Open(path)
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

// 利用bufio这个包读取文件
func readFromFileByfio(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("open file err %v\n", err)
		return
	}
	//记得关闭文件
	defer file.Close()
	//创建读内容的对象
	read := bufio.NewReader(file)
	
}
