package main

import (
	"fmt"
	"net"
)

// tcp  server端

func main() {
	// 1.本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.0:20000")
	if err != nil {
		fmt.Printf("start server no 127.0.0.0:20000 failed,err:%v", err)
		return
	}
	// 2.等待别人建立连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("accept failed,err:%v", err)
		return
	}
	// 3.进行通信
	var temp [128]byte
	n, err := conn.Read(temp[:])
	if err != nil {
		fmt.Printf("read  from conn failed,err:%v", err)
		return
	}
	fmt.Println(string(temp[:]))
}
