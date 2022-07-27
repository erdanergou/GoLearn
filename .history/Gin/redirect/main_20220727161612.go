package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 重定向

func main() {
	// 创建路由
	eng := gin.Default()

	eng.GET("/redirect", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")

	})

	eng.Run()
}
