package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// URI数据解析与绑定

type User struct {
	//binding:"required" 修饰的字段为必选字段，若接受时为空，则报错
	Username string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	// 创建一个路由
	eng := gin.Default()
	eng.GET("/:username/:password", func(ctx *gin.Context) {
		// 声明接收信息的结构体
		var user User
		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type 自动推断
		err := ctx.ShouldBindUri(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		if user.Username != "root" || user.Password != "root" {
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
