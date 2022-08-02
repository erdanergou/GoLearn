package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin hello world

// 当我们渲染的HTML文件中引用了静态文件时，我们需要在渲染页面前调用gin.Static方法即可。

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()   也可以创建不适用中间件的  eng:=gin.New()
	eng := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context封装了request和response
	eng.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world!")
	})
	eng.POST("/xxxPost", getting)
	// 3.监听端口,默认在8080端口
	eng.Run()
}

func getting(c *gin.Context) {

}
