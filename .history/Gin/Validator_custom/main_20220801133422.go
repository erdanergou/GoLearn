package main

// 自定义验证

type Login struct {
	Username string `uri:"username" validate:"checkName"`
	password string `uri:"password"`
}

func main() {

}

func checkName(){
	
}
