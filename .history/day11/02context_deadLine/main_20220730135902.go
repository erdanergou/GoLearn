package main

import (
	"context"
	"time"
)

//context.WithDeadline

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)


	//	尽管穿拖鞋、
}