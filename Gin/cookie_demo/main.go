package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value, err := ctx.Cookie("user")
		fmt.Printf("user is %#v\n", value)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "don't have this user"})
			ctx.Abort()
			return
		}
		ctx.Next()
		ctx.JSON(200, gin.H{"msg": value})

	}
}

func main() {
	//创建路由
	eng := gin.Default()

	eng.GET("/login", loginFunc)
	eng.GET("/home", MiddleWare(), homeFunc)
	eng.Run()
}
func loginFunc(ctx *gin.Context) {
	// 登录
	username := ctx.Query("username")
	password := ctx.Query("password")
	if username != "root" || password != "root" {
		ctx.Redirect(404, "localhost:8080/login")
		// 如果验证不通过，不在调用后续的函数处理

		return
	}
	ctx.SetCookie("user", "true", 60, "/", "localhost", false, true)
}

func homeFunc(ctx *gin.Context) {
	// 首页
	ctx.JSON(200, gin.H{"data": "home"})
}
