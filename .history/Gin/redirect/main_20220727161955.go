package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 重定向
// Redirect

func main() {
	// 创建路由
	eng := gin.Default()

	eng.GET("/redirect", func(ctx *gin.Context) {
		// 支持内部和外部重定向
		ctx.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")

	})

	eng.Run()
}
