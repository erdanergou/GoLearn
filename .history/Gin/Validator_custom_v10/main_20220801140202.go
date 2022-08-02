package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

/*


Var 方法使用 tag（标记）验证方式验证单个变量。

func (*validator.Validate).Var(field interface{}, tag string) error
*/

func main() {
	validate := validator.New()
	// 验证变量
	email := "admin#admin.com"
	// email := ""
	err := validate.Var(email, "required,email")
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		fmt.Println(validationErrors)
		// output: Key: '' Error:Field validation for '' failed on the 'email' tag
		// output: Key: '' Error:Field validation for '' failed on the 'required' tag
		return
	}
}
