package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//打开文件写内容

func main() {
	writeDemo4()
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

// 在go语言中一个汉字3个字节，英文字母1字节
// 向文件中间插入内容
func writeDemo4() {
	str := "小垃圾在坚持"
	// 打开待操作文件
	file, err := os.OpenFile("./xx.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file falied %v", err)
		return
	}
	defer file.Close()
	// 因为无法直接在文件中间插入内容，所有要借助临时文件
	fileobj, err := os.OpenFile("./sb.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("create file falied %v", err)
		return
	}
	defer fileobj.Close()

	//读取源文件写入临时文件
	var ret [16]byte
	n, err := fileobj.Read(ret[:])
	if err != nil {
		fmt.Printf("read file falied %v", err)
		return
	}
	fmt.Printf("读了%d个字符\n", n)
	// 写入临时文件中
	fileobj.Write([]byte(str))

	// 把源文件中的内容写入临时文件
	var x [1024]byte
	for {
		n, err := file.Read(x[:]); 
		if err != nil; 
	}
	fmt.Println(string(ret[:]))
	// 在window下每行结束有/r/n，占两字节
	fileobj.Seek(16, 0) // 光标移动

}
