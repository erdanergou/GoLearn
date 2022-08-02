package main

// 参数验证

// 结构体验证
type Person struct {
	Name string `form:"name" binding:"requrie"`
	//不能为空并且大于10
	Age int `form:"age" binding:"requrie,gt=10"`
}

func main() {
	ctx := gin.de
}
