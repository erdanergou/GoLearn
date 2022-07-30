package main

import (
	"context"
	"time"
)

//context.WithDeadline

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)


	//	尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践
	// 如果不这么做，
}
