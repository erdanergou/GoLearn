package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子操作
// 针对整数数据类型（int32、uint32、int64、uint64）我们还可以使用原子操作来保证并发安全，通常直接使用原子操作比使用锁操作效率更高。

var x int64 = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	// lock.Lock()
	// x++
	// lock.Unlock()
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	// wg.Add(1000)
	// for i := 0; i < 1000; i++ {
	// 	go add()
	// }
	// wg.Wait()
	// fmt.Println(x)

	ok := atomic.Comp
}
