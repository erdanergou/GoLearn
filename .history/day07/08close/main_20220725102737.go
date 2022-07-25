package main

import "fmt"

// close
/*
当对一个关闭了的通道进行取值时，是能进行取值的，
	x, ok := <-ch1 
*/


func main() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	close(ch1)
	// for x := range ch1 {
	// 	fmt.Println(x)
	// }
	<-ch1
	<-ch1
	x, ok := <-ch1
	fmt.Println(x, ok)
}
