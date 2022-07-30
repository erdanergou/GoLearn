package main

import (
	"context"
	"time"
)

//context.WithDeadline

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	context.WithDeadline(context.Background(),d)
}
