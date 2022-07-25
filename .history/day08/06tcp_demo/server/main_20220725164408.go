package main

import (
	"fmt"
	"net"
)

// tcp  server端

func main() {
	// 1.本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.0:20000")
	if err != nil{
		fmt.Printf("start server no 127.0.0.0 ")
	}
	// 2.等待别人建立连接
	// 3.进行通信
}
