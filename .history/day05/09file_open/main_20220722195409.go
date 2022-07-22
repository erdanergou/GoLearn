package main

import (
	"fmt"
	"os"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

//打开文件

func main(){
	file,err := os.Open("./main.go")
	if err != nil{
		fmt.Printf("open file err %s\n",err)
		return
	}
	//记得关闭文件
	def file.Clos()
}