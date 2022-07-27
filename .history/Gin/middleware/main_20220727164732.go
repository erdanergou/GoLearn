package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

/*gin中间件
gin可以构建中间件，但它只对注册过的路由函数起作用
对于分组路由，嵌套使用中间件，可以限定中间件的作用范围
中间件分为全局中间件，单个路由中间件和群组中渐渐
gin中间件必须是一个gin.HandlerFunc类型
*/

// 1.全局中间件

//定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		befor := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到context中，可以通过Get取
		ctx.Set("req", "中间件")
		// 执行中间件
		ctx.Next()
		status := ctx.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		now := time.Since(befor) // 是time.Now().Sub()的缩写
		fmt.Println("time: ", now)
	}
}

func main() {
	//创建路由
	eng := gin.Default()
	// 注册中间件
	eng.Use(MiddleWare())
	// {}为了代码规范
	{
		eng.GET("/middleware", func(ctx *gin.Context) {
			//取值
			value, exists := ctx.Get("req")
			if !exists {
				fmt.Println("req not exists")
			}
			fmt.Println("request: ", value)
			// 页面接收
			ctx.JSON(200,gin.H{"req":})

		})
	}
	eng.Run()
}
