package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitGroup

func fRand() {
	rand.Seed(time.Now().UnixNano()) // 生成随机数要加随机数种子，否则生成的随机数将一致
	for i := 0; i < 5; i++ {
		r := rand.Intn(10) // 0<=x<10
		fmt.Println(r)
	}
}
func f1(i int) {
	defer wg.Done()  // 当函数
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300))) //随机睡眠时长
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {
	fRand()
	for i := 0; i < 10; i++ {
		wg.Add(1) // 每次给wg加1
		go f1(i)
	}

	// 如何知道这10个goroutine都结束了
	wg.Wait()
}
