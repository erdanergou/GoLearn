package main

import (
	"log"
	"os"
	"time"
)

func main() {
	os.OpenFile("./xx.log",os.O_APPEND|os.O_RDWR)
	for {
		log.Println("这是一个测试的日志")
		time.Sleep(time.Second * 3)
	}
}
