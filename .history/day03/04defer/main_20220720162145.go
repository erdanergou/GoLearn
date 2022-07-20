package main

import "fmt"

//defer

func deferDemo(){
	fmt.Println("嘿嘿嘿")
	defer fmt.Fprintln()
}

func main(){

}