package main

import (
	"fmt"
	"math/rand"
)

// waitGroup

func hello() {
	// rand.Seed()
	for i := 0; i < 5; i++ {
		r := rand.Intn(10) // 0<=x<10
		fmt.Println(r)
	}
}

func main() {
	hello()
}
