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
	k := v.Kind() //值得种类
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

// 通过反射设置变量的值
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		// 反射中使用Elem()方法获取指针对应的值
		v.Elem().SetInt(200) 
	}
}
func main() {

	var a float32 = 1.234
	reflectType(a)

	var b int64 = 3
	reflectType(b)

	var c cat
	reflectType(c)

	var d float32 = 3.14
	reflectValue(d)

	var e float64 = 3.123
	reflectValue(e)

	reflectValue(c)

	// reflectSetValue(&b)
	reflectSetValue2(&b)
	fmt.Println(b)
}
