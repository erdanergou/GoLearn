package main

import (
	"fmt"
	"reflect"
)

// reflect


func reflectType(x interface{}){
	v := reflect.ValueOf(x)
	fmt.Printf("type")
}

func main() {
	var a float32 = 1.234
	
}