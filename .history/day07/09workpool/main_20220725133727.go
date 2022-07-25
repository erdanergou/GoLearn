package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
	wg.Done()
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// 五个任务
	go func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
		}
		close(jobs)
	}()
	// 输出结果
	go func() {
		for a := 1; a <= 5; a++ {
			<-results
		}
	}()
	// 开启三个goroutine
	wg.Add(3)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	wg.Wait()

}
