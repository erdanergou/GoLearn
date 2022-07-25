package main

import "net"

// tcp  server端

func main() {
	// 1.本地端口启动服务
	net.Listen("tcp","127.")
	// 2.等待别人建立连接
	// 3.进行通信
}
