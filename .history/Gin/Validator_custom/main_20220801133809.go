package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 自定义验证

/*
   对绑定解析到结构体上的参数，自定义验证功能
   比如我们需要对URL的接受参数进行判断，判断用户名是否为root如果是root通过否则返回false
*/
type Login struct {
	Username string `uri:"username" validate:"checkName"`
	password string `uri:"password"`
}

func main() {
	eng := gin.Default()

	validate := validator.New()
	
	eng.GET("/:username/:password",func(ctx *gin.Context) {
		var login Login
		 //注册自定义函数，与struct tag关联起来
	})

	eng.Run()
}

func checkName(fl validator.FieldLevel) bool {
	if fl.Field().String() != "root" {
		return false
	}
	return true
}
