package main

import (
	"fmt"
	"reflect"
)

// reflect

type cat struct {
}

// TypeOf获取到类型部分
// type类型，kind种类
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type name:%v type kind %v\n", v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整形的原始值，如何通过int64()强制类型转换
		fmt.Printf("type is int64,value is %d\n", int64(v.Int()))
	case reflect.Float64:
		// v.Float()从反射中获取整形的原始值，如何通过float64()强制类型转换
		fmt.Printf("type is float64,value is %f\n", float64(v.Float()))
	case reflect.Float32:
		// v.Float()从反射中获取整形的原始值，如何通过float32()强制类型转换
		fmt.Printf("type is float32,value is %f\n", float32(v.Float()))

	}
}

func main() {

	
	var a float32 = 1.234
	reflectType(a)

	var b int64 = 3
	reflectType(b)

	var c cat
	reflectType(c)
}
