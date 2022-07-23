package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//打开文件

func main() {
	// readFromFile1("./main.go")
	// readFromFileByfio("./main.go")
	readFromFileByIoutil("./main.go")
}

func readFromFile1(path string) {
	file, err := os.Open(path)
	// defer file.Close() //不能写在此处，因为若有err出现，则file为空，无法调用close方法
	//记得关闭文件
	defer file.Close()
	if err != nil {
		fmt.Printf("open file err %v\n", err)
		return
	}

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
	for {

		line, err := read.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed,err:%v\n", err)
			return
		}
		fmt.Println(line)
	}

}

func readFromFileByIoutil(path string) {
	ret, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("open file err %v\n", err)
		return
	}
	fmt.Println(string(ret))

}
