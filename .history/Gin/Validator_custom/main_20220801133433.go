package main

import "github.com/go-playground/validator/v10"

// 自定义验证

type Login struct {
	Username string `uri:"username" validate:"checkName"`
	password string `uri:"password"`
}

func main() {

}

func checkName(validator.FieldError){

}
