package main

import (
	"context"
	"time"
)

//context.WithDeadline

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx,c context.WithDeadline(context.Background(),d)
}
