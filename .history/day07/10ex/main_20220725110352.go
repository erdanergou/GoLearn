package main

// job
type Job struct {
	x int16
}

// result
type Result struct {
	Job
	result int64
}

var jobChan = make(chan *Job, 100)
var resultChan = make(chan *Result, 100)

func sum(ch1 <-chan *Job){
	
}

func main() {

}
