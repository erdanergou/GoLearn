package main

import "github.com/gin-gonic/gin"

// 路由组
/*
	将拥有共同URL前缀的路由划分为一个路由组。习惯性一对{}包裹同组的路由
*/

// 路由原理
/*
httproute 会将所有路由规则构造成一颗前缀树

*/

func main() {

	// 1.创建路由
	// 默认使用两个中间件
	eng := gin.Default()

	// 路由组 处理Get请求
	v1 := eng.Group("/v1")
	// {}是一个书写规范
	{
		v1.GET("/login", login)
		v1.GET("/submint", submit)
	}

	v2 := eng.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}
	eng.Run()
}

func login(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "login")
	ctx.String(200, name)
}

func submit(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "submit")
	ctx.String(200, name)
}
