package main

import (
	"fmt"
	"strconv"
	"sync"
)

// go语言中内置的map并不是并发安全的
// 超过20时会报错

/*
像这种场景下就需要为 map 加锁来保证并发的安全性了，Go语言的sync包中提供了一个开箱即用的并发安全版 map——sync.Map。
开箱即用表示其不用像内置的 map 一样使用 make 函数初始化就能直接使用。
*/
var m = make(map[string]int)
var lock sync.Mutex

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

// func main() {
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 21; i++ {
// 		wg.Add(1)
// 		go func(n int) {
// 			key := strconv.Itoa(n)
// 			lock.Lock()
// 			set(key, n)
// 			lock.Unlock()
// 			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }

var m2 sync.Map{}
func main() {
	
}
