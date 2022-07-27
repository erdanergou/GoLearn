package main

import "github.com/gin-gonic/gin"

// 同步和异步
/*
goroutine机制可以方便的实现同步异步
另外，在启动新的goroutine时，不应该使用原始上下文，必须使用他的只读副本
*/

func main() {
	// 创建路由
	eng := gin.Default()

	// 异步
	eng.GET("long_async",func(ctx *gin.Context) {
		// 需要搞一个副本
		copyctx.Copy()
	})


	eng.Run()
}
