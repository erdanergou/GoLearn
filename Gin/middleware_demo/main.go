package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func MiddleWare(c *gin.Context) {
	befor := time.Now()
	c.Next()
	timespace := time.Since(befor)
	fmt.Printf("程序用时：%v\n", timespace)
}

func main() {
	// 创建路由
	eng := gin.Default()
	// 注册中间件
	eng.Use(MiddleWare)
	g1 := eng.Group("/shopping")
	{

		g1.GET("/index", func(ctx *gin.Context) {
			time.Sleep(3 * time.Second)
			ctx.JSON(200, gin.H{"name": "index"})
		})
		g1.GET("/home", func(ctx *gin.Context) {
			time.Sleep(3 * time.Second)
			ctx.JSON(200, gin.H{"name": "home"})
		})

	}
	eng.Run()
}
