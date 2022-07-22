package main

import (
	"fmt"
	"os"
)

//打开文件写内容

func main() {
	_, err := os.OpenFile("./xx.txt", os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file err %v\n", err)
		return
	}

}
