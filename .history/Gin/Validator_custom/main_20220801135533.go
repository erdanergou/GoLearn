package main

import (
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 自定义验证

/*
   对绑定解析到结构体上的参数，自定义验证功能
   比如我们需要对URL的接受参数进行判断，判断用户名是否为root如果是root通过否则返回false
*/
type Login struct {
	Username string `uri:"username" validate:"checkName"`
	Password string `uri:"password"`
}

// Booking contains binded and validated data.
type Booking struct {
	//定义一个预约的时间大于今天的时间
	CheckIn time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	//gtfield=CheckIn退出的时间大于预约的时间
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func main() {
	eng := gin.Default()

	// validate := validator.New()

	// eng.GET("/:username/:password", func(ctx *gin.Context) {
	// 	var login Login
	// 	//注册自定义函数，与struct tag关联起来
	// 	err := validate.RegisterValidation("checkName", checkName)
	// 	err = ctx.ShouldBindUri(&login)
	// 	if err != nil {
	// 		ctx.JSON(500, gin.H{"err": err})
	// 		return
	// 	}
	// 	err = validate.Struct(login)
	// 	if err != nil {
	// 		for _, err := range err.(validator.ValidationErrors) {
	// 			fmt.Println(err)
	// 		}
	// 		return
	// 	}
	// 	fmt.Printf("%#v\n", login)
	// 	fmt.Println("success")
	// })

	//注册验证
	v, ok := binding.Validator.Engine().(*validator.Validate)
	validate := validator.New()
	validate.
	if ok {
		//绑定第一个参数是验证的函数第二个参数是自定义的验证函数
		v.RegisterValidation("bookabledate", bookableDate)
	}

	eng.GET("/5lmh", getBookable)

	eng.Run()
}

func checkName(fl validator.FieldLevel) bool {
	if fl.Field().String() != "root" {
		return false
	}
	return true
}

func bookableDate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	//field.Interface().(time.Time)获取参数值并且转换为时间格式
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Unix() > date.Unix() {
			return false
		}
	}
	return true
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
