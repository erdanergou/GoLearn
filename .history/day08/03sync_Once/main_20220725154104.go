package main

// sync.Once
// Go语言中的sync包中提供了一个针对只执行一次场景的解决方案

var wg sync.


func f1(ch1 chan<- int) {
	defer wg.Done()
}

func main() {

}
