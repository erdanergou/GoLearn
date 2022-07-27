package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin数据解析和绑定
/*
1.json数据解析和绑定
	客户端传参，后端接收并解析到结构体



*/

// 定义接受数据的结构体
type User struct {
	//binding:"required" 修饰的字段为必选字段，若接受时为空，则报错
	Username string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

/*cmd命令窗口进行测试
curl localhost:8080/loginJson -H 'content-type:application/json' -d "{\"username\":\"root\",\"password\":\"root1\"}" -X POST

*/

func main() {
	// 创建路由
	eng := gin.Default()

	// Json绑定
	eng.POST("loginJson", func(ctx *gin.Context) {
		// 声明接收的变量
		var u1 User
		// 将request的Body中的数据，自动按照json格式解析到结构体
		err := ctx.ShouldBindJSON(&u1)
		if err != nil {
			// 返回错误信息
			// gin.H 封装了生成json数据的工具
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// 判断用户名密码是否正确
		if u1.Username != "root" || u1.Password != "root" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": "304",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": "200",
		})
	})
	eng.Run()
}
