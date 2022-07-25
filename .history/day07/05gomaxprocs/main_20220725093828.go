package main

import (
	"fmt"
	"runtime"
	"sync"
)
/*
goroutime调动模型
	GMP：G表示goroutine，M和操纵系统线程做映射的关系，P是调度者，管理goroutine

	M:N 把M个goroutine分配个N个操作系统线程去执行

	goroutine初始栈的大小是2k

*/



// GOMAXPROCS
var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

func main() {
	runtime.GOMAXPROCS(1) //不进行设置时，默认为CPU的逻辑核心数，默认跑满整个CPU
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
