package main

import "fmt"

//panic 和 recover

func funcA(){
	fmt.Println("a")
}
func funcB(){
	fmt.Println("b")
}
func funcA(){
	fmt.Println("c")
}

func main(){

}