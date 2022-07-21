package main

import "fmt"

//结构体模拟实现其他语言中的继承

type animal struct{
	name string
}

//给动物一个移动的方法
func (a animal)move(){
	fmt.Printf("%s 会动",a.name)
}


type dog struct{
	feet uint8
}

//
func main() {

}
