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
		fmt.Println("open fi")
	}
}