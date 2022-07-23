package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//打开文件写内容

func main() {
	writeDemo2()
}
func writeDemo1() {
	file, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file err %v\n", err)
		return
	}
	defer file.Close()
	// write
	file.Write([]byte("zheng bu hui le!"))
	// write string
	file.WriteString("整不会了")

}

func writeDemo2() {
	file, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file err %v\n", err)
		return
	}
	defer file.Close()
	// 创建一个写对象
	write := bufio.NewWriter(file)
	write.WriteString("小垃圾学Go") //写到缓存中，需要推送
	write.Flush()               //将缓存中的内容写入文件

}

func writeDemo3() {
	str := "小垃圾不想学了"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("write file failed,err:%v\n", err)
		return
	}

}

// 向文件中间插入内容
func writeDemo4(){
	str := "小垃圾在坚持"
	file,err := os.OpenFile("./xx.txt",os.O_RDWR,0644)
	if err!=nil{
		fmt.Printf("open file falied %v",err)
		return
	}
	file.Seek(1,0)
}