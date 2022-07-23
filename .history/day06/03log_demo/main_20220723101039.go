package main

import (
	"log"
	"os"
	"time"
)

func main() {
	os.OpenFile()
	for {
		log.Println("这是一个测试的日志")
		time.Sleep(time.Second * 3)
	}
}
