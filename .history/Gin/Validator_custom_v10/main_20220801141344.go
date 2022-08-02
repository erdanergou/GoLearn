package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
)

/*变量验证

Var 方法使用 tag（标记）验证方式验证单个变量。

func (*validator.Validate).Var(field interface{}, tag string) error
*/

/*结构体验证

结构体验证结构体公开的字段，并自动验证嵌套结构体，除非另有说明。

func (*validator.Validate).Struct(s interface{}) error

*/

func main() {

	// validate := validator.New()
	// // 验证变量
	// email := "admin#admin.com" // admin@admin.com
	// // email := ""
	// err := validate.Var(email, "required,email")
	// if err != nil {
	// 	validationErrors := err.(validator.ValidationErrors)
	// 	fmt.Println(validationErrors)
	// 	// output: Key: '' Error:Field validation for '' failed on the 'email' tag
	// 	// output: Key: '' Error:Field validation for '' failed on the 'required' tag
	// 	return
	// }

	validate := validator.New()
	type User struct {
		ID     int64  `json:"id" validate:"gt=0"`
		Name   string `json:"name" validate:"required"`
		Gender string `json:"gender" validate:"required,oneof=man woman"`
		Age    uint8  `json:"age" validate:"required,gte=0,lte=130"`
		Email  string `json:"email" validate:"required,email"`
	}
	user := &User{
		ID:     1,
		Name:   "frank",
		Gender: "boy",
		Age:    180,
		Email:  "gopher@88.com",
	}
	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		// output: Key: 'User.Age' Error:Field validation for 'Age' failed on the 'lte' tag
		// fmt.Println(validationErrors)
		fmt.Println(validationErrors)
		return
	}
	
}
