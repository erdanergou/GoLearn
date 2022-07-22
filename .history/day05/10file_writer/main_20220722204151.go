package main

import (
	"fmt"
	"os"
)

//打开文件写内容

func main() {
	file, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file err %v\n", err)
		return
	}
	// write
	file.Write([]byte("zheng bu hui le!"))
	// write string
	file.WriteString("整不会了")
	
}
