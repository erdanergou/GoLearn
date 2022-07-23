package main

import (
	"log"
	"os"
	"time"
)

func main() {
	fileObj,erros.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	for {
		log.Println("这是一个测试的日志")
		time.Sleep(time.Second * 3)
	}
}
