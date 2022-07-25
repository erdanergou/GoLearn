package main



type Job struct{
	x int16
}

type Result struct{
	Job
	result int64
}

var jobChan = make(chan, 0)
func main() {

}
