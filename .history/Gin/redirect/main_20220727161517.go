package main

import "github.com/gin-gonic/gin"

// 重定向

func main() {
	// 创建路由
	eng := gin.Default()

	eng.GET("redirect",func(ctx *gin.Context) {
		

	})


	eng.Run()
}
