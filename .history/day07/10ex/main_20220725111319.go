package main

import (
	"math/rand"
	"time"
)

// job
type Job struct {
	x int64
}

// result
type Result struct {
	Job
	result int64
}

var jobChan = make(chan *Job, 100)
var resultChan = make(chan *Result, 100)

func createInt64(j chan<- *Job) {
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
	for {
		job := <-j
		sum := int64(0)
		n := (*job).x
		for n > 0 {
			sum += n % 10

		}
	}
}

func main() {

}
