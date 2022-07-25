package main

import (
	"fmt"
	"net"
)

// tcp client

func main() {
	// 与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.0 20000")
	if err != nil {
		fmt.Println("dial")
	}
	// 发送数据

}
