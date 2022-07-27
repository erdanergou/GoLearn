package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建路由
	eng := gin.Default()
	// 服务端要给客户端cookie
	eng.GET("/cookie", func(ctx *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := ctx.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			//给客户端设置cookie
			// maxAge int 单位为秒
			// , path,  cookie所在目录
			// domain string,  域名
			// secure 是否只能通过http访问
			// httpOnly bool 是否允许别人通过js获取自己的cookie
			ctx.SetCookie("key_cookie", "zhangsan", 60, "/", "localhost", false, true)
		}
		fmt.Printf("cookie的值是：%s\n", cookie)
	})
	eng.Run()
}
