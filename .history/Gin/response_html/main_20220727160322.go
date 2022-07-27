package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

//多种响应方式

func main() {
	// 创建路由
	eng := gin.Default()
	// 1、Json
	eng.GET("/someJson", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "someJson",
			"status":  200,
		})
	})

	// 2.结构体响应
	eng.GET("/someStruct", func(ctx *gin.Context) {
		var Msg struct {
			Name    string
			Message string
			Number  int
		}
		Msg.Name = "root"
		Msg.Message = "message"
		Msg.Number = 20
		ctx.JSON(200, Msg)
	})

	// 3.XMl响应
	eng.GET("/someXml", func(ctx *gin.Context) {
		ctx.XML(200, gin.H{"message": "abc"})
	})

	// 4.Yaml响应
	eng.GET("/someYaml", func(ctx *gin.Context) {
		ctx.YAML(200, gin.H{"name": "zhangsan"})
	})

	// 5.protobuf 格式，谷歌开发的高效存储读取的工具
	eng.GET("/smoeProtoBuf", func(ctx *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		// 定义数据
		label := "label"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		ctx.ProtoBuf(200, data)
	})
	eng.Run()
}
