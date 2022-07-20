package main

import "fmt"

func f1(s1 string) {
	fmt.Println(s1)
}

func f2(a, b int) int {
	return a + b
}
func f3(a int , b ... int){
	
}

func main() {
	f1("asd")
	a := f2(3, 4)
	fmt.Println(a)
}
