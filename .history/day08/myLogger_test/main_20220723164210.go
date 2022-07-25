package main

import "GoLearn/day06/04myLogger"


var log myLogger.Logger  // 声明一个接口


// 测试自建日志库
func main() {
	// logger := myLogger.NewConsoleLog("Error")
	logger := myLogger.NewFileLogger("Info", "./", "test.log", 10*1024)


	// n := 0
	for {
		logger.Debug("这是一条Debug日志")
		logger.Info("这是一条Info日志")
		logger.Warning("这是一条Warning日志")
		id := 1
		str := "发送了垃圾"
		logger.Error("这是一条Error日志,id:%d,str:%s", id, str)
		logger.Fatal("这是一条Fatal日志")
		// n++
		// if n >= 10{
		// 	break
		// }
	}
}
