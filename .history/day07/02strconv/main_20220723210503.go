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
	strconv.Atoi(str)

}
