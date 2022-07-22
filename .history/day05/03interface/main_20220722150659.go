package main

import "fmt"

// 接口
// 接口是一种类型


type cat struct{}

type dog struct{}

func (c cat)speak(){
	fmt.Println("喵喵喵~~~")
}

func (d dog)speak(){
fmt.Println("汪汪汪")
}

func main() {

}
