package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 同步和异步
/*
goroutine机制可以方便的实现同步异步
另外，在启动新的goroutine时，不应该使用原始上下文，必须使用他的只读副本
*/

func main() {
	// 创建路由
	eng := gin.Default()

	// 异步
	eng.GET("long_async", func(ctx *gin.Context) {
		// 需要搞一个副本
		copyContext := ctx.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			fmt.Printf("异步执行:" + copyContext.Request.URL.Path)
		}()
	})


	
	eng.Run()
}
