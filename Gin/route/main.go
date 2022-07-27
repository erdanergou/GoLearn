package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin 路由
/*
	基本路由
	gin框架中采用的的路由库是基于httprouter做的

	restful风格的api
	gin支持restful风格的api即Representational State Transfer 的缩写，是一种互联网应用程序api设计理念：URL定位资源，用HTTP描述操作

	路由eg:
	1.获取文章 /blog/getXxx   GET   blog/Xxx
	2.添加 /blog/addXxx		  POST  blog/Xxx
	3.修改 /blog/updateXxx	  PUT   blog/Xxx
 	4.删除 /blog/delXxx		  DELETE blog/Xxx

	API参数
	可以通过Context的Param方法获取API参数
	localhost:8080/xxx/zhangsan

	URL参数
	URL参数可以通过DefaultQuery() 和 Query()方法获取
	DefaultQuery()若参数不对则返回默认值，Query()若不存在则返回空串
	Api?name="zs"&


	表单参数
	表达传递的为post请求，http常见的传输格式为四种
	application/json
	application/x-www-form-urlencoded
	application/xml
	multipart/from-data

	表单参数可以通过PostForm()方法获取，该方法默认解析的是x-www-form-urlencoded或from-data格式的参数

	上传单个文件
	multipart/from-data格式用于文件上传
	gin文件上传与原生的net/http方法类似，不同在于gin把原生的request封装到ctx.Request中
	将form表单的enctype修改为"multipart/form-data"且input type改为file
	后台使用FormFile进行接收

	上传多个文件
	需要将input添加multiple属性 后台使用MultipartForm进行获取

*/

func main() {
	// 1.创建路由
	// 默认使用两个中间件
	eng := gin.Default()
	// 限制表单上传大小，默认为32mb
	// eng.MaxMultipartMemory = 8 << 20

	// api参数 :
	eng.GET("/user/:name/*action", func(ctx *gin.Context) {
		value := ctx.Param("name")
		action := ctx.Param("action")
		ctx.String(http.StatusOK, value+" is "+action)
	})

	//url参数
	eng.GET("/welcome", func(ctx *gin.Context) {
		// 第二个参数是默认值
		name := ctx.DefaultQuery("name", "jack")
		ctx.String(http.StatusOK, name)
	})

	// 表单参数
	eng.POST("/test/login", func(ctx *gin.Context) {
		// 表单参数设置默认值
		sex := ctx.DefaultPostForm("sex", "maele")
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		// 多选框
		hobby := ctx.PostFormArray("hobby")
		ctx.String(http.StatusOK,
			fmt.Sprintf("username:%s,password:%s,sex:%s,hobby:%v\n", username, password, sex, hobby))
	})

	// 表单参数上传文件
	eng.POST("/test/uploadFile", func(ctx *gin.Context) {
		// 表单取单个文件
		file, _ := ctx.FormFile("file")
		// 表单取多个文件
		// ctx.MultipartForm()
		log.Println(file.Filename)
		// 传到项目根目录
		ctx.SaveUploadedFile(file, file.Filename)
		// 打印信息
		ctx.String(200, fmt.Sprintf("%s upload!", file.Filename))
	})
	eng.Run(":8000")
}
