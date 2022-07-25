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

func sum(ch1 <-chan<- *Job) {
	// 循环生产int64类型的随机数，发送到jobChan
	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < 5; i++ {
		x := rand.Int63()
		newJob := &Job{
			x: x,
		}
		ch1 <- newJob
	}
}

func main() {

}
