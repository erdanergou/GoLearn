package main

import (
	"math/rand"
	"sync"
	"time"
)

// job
type Job struct {
	x int64
}

// result
type Result struct {
	job    *Job
	result int64
}

var jobChan = make(chan *Job, 100)
var resultChan = make(chan *Result, 100)
var wg sync.WaitGroup

func createInt64(j chan<- *Job) {
	wg.Done()
	// 循环生产int64类型的随机数，发送到jobChan
	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < 5; i++ {
		x := rand.Int63()
		newJob := &Job{
			x: x,
		}
		j <- newJob
	}
}

func sendInt64(j <-chan *Job, r chan<- *Result) {
	// 从jobChan中取出各随机数并计算各位数的和，将结果发送到resultChan中
	wg.Done()
	for {
		job := <-j
		sum := int64(0)
		n := (*job).x
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &Result{
			job:    job,
			result: sum,
		}
		r <- newResult
	}
}

func main() {
	wg.Add()
	go createInt64(jobChan)
	for i := 0; i < 3; i++ {
		go sendInt64(resultChan)
	}
}
