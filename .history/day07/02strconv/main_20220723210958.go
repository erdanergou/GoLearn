package main

import (
	"fmt"
	"strconv"
)

//strconv

func main() {
	// i := int32(100)
	// ret2 := string(i)
	// fmt.Println(ret2)

	str := "190"
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Printf("parse int errr,%v", err)
	}
	fmt.Printf("%T\n", ret1)
	fmt.Println(ret1)

	//
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v\n", retInt)

	ret3 := strconv.Itoa(100)
	fmt.Printf("%#v\n", ret3)

	boolstr := "true"
	retbool, _ := strconv.ParseBool(boolstr)
	fmt.Printf("%#v\n %T", retbool, retbool)

	floatstr := "1.234"
	retfloat, _ := strconv.ParseFloat(floatstr, 64)
	
}
