package main

import (
	"fmt"
	"time"
)

// 时间

func main() {
	now := time.Now()
	// 完整时间，eg：2022-07-23 09:10:42.4661592 +0800 CST m=+0.002036801
	fmt.Println(now)
	// 时间戳
	fmt.Println(now.Unix())
	// 纳秒时间戳
	fmt.Println(now.UnixNano())
	// 将时间戳转为时间
	t := now.Unix()
	
}
