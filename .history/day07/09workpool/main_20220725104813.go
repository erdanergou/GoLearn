package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d", id, j)
		results <- j * 2
	}
}

func main() {
	
}
