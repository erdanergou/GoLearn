package main

import "github.com/gin-gonic/gin"

/*gin中间件
gin可以构建中间件，但它只对注册过的路由函数起作用
对于分组路由，嵌套使用中间件，可以限定中间件的作用范围
中间件分为全局中间件，单个路由中间件和群组中渐渐
gin中间件必须是一个gin.HandlerFunc类型
*/

// 1.全局中间件

func main() {
	//
	eng := gin.Default()

	eng.Run()
}
