package main

import (
	"fmt"
	"reflect"
)

// reflect
// TypeOf获取到类型部分
// type类型，kind种类



type cat struct {
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type name:%v type kind %v\n", v.Name(), v.Kind())
}

func main() {
	var a float32 = 1.234
	reflectType(a)

	var b int64 = 3
	reflectType(b)

	var c cat
	reflectType(c)
}
