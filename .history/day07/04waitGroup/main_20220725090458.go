package main

import (
	"fmt"
	"math/rand"
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
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt
}

func main() {
	fRand()
	for i := 0; i < 10; i++ {
		go f1()
	}
}
