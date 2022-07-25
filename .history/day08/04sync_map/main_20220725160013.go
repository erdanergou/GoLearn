package main

import (
	"fmt"
	"strconv"
	"sync"
)

// go语言中内置的map并不是并发安全的
var m = make(map[string]int)
var lock sync.Mutex

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			lock.lock()
			set(key, n)
			
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
