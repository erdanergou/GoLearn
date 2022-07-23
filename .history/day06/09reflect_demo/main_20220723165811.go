package main

import (
	"fmt"
	"reflect"
)

// reflect

func reflectType(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("type:%v\n", v)
}

func main() {
	var a float32 = 1.234
	reflectType(a)

	var b int64 = 3
	reflectType(b)

}
