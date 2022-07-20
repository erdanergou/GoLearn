package main

import "fmt"

//panic 和 recover

func funcA(){
	fmt.Println("a")
}
func funcB(){
	fmt.Println("b")
}
func funcC(){
	fmt.Println("c")
}

func main(){
	funcA()
	funcB()
	funcC()
}