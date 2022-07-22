package main

import "os"

//打开文件写内容

func main() {
	fileos.OpenFile("./xx.txt",os.O_APPEND,0644)

}
