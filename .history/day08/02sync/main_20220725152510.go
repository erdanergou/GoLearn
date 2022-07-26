package main

import (
	"fmt"
	"sync"
	"time"
)

//锁
var x int = 0
var wg sync.WaitGroup
var lock sync.Mutex // 互斥锁
var rw sync.RWMutex

/*
读写互斥锁分为读锁与写锁（读的次数远远大于写）
当一个 goroutine 获取到读锁之后，其他的 goroutine 如果是获取读锁会继续获得锁
如果是获取写锁就会等待；而当一个 goroutine 获取写锁之后，其他的 goroutine 无论是获取读锁还是写锁都会等待。

*/

// 互斥锁
func add() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
}

// 读写锁
func read() {
	defer wg.Done()
	// lock.Lock()
	rw.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	// lock.Unlock()
}

func write() {
	defer wg.Done()
	lock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	lock.Unlock()
}

func main() {
	//互斥锁
	// wg.Add(2)
	// go add()
	// go add()
	// wg.Wait()
	// fmt.Println(x)

	// 读写锁
	now := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(now))
}