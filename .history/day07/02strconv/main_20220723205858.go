package main

import (
	"errors"
	"fmt"
	"strconv"
)

//strconv

func main() {
	// i := int32(100)
	// ret2 := string(i)
	// fmt.Println(ret2)

	str:="190"
	ret1,err := strconv.ParseInt(str,10,64)
	if err!=nil{
		fmt.Printf("parse int errr,%v",err)
	}
	

}
