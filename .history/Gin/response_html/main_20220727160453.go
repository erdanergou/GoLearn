package main

import (
	"github.com/gin-gonic/gin"
)

// html渲染

func main() {
	// 创建路由
	eng := gin.Default()
	// 加载模板文件
	eng.LoadHTMLGlob("templates/*")
	// eng.LoadHTMLFiles("templates/index.tmpl")

	eng.GET("/index",)
	eng.Run()
}
