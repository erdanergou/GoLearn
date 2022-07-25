package main

import "sync"

// channel练习

var ch1 chan int
var ch2 chan int
var wg sync.


func f1(ch chan int) {
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
}

func main() {

}
