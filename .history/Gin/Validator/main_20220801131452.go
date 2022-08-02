package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 参数验证

// 结构体验证
type Person struct {
	Name string `form:"name" binding:"requrie"`
	//不能为空并且大于10
	Age int `form:"age" binding:"requrie,gt=10"`
}

func main() {
	ctx := gin.Default()
	ctx.GET("/test", func(ctx *gin.Context) {
		var person Person
		if err := ctx.ShouldBindJSON(&person); err != nil {
			// ctx.String(500, fmt.Sprint(err))
			fmt.Printf("%#v\n", person)
			ctx.JSON(500, gin.H{"err": err})
			return
		}
		fmt.Printf("%#v\n", person)
		ctx.JSON(200, gin.H{"data": person})
	})
	ctx.Run()
}
