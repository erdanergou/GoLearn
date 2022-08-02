package main

// 参数验证


// 结构体验证
type Person struct{
	Name string `form:"name" binding:"requrie"`
	Age string `form:"age" binding:"requrie"`
}


func main() {

}	
