package main

import (
	"fmt"
	"math/rand"
	"time"
)

// waitGroup

func fRand() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 5; i++ {
		r := rand.Intn(10) // 0<=x<10
		fmt.Println(r)
	}
}

func main() {
	fRand()
}
