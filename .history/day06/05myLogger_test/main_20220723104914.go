package main

// import "GoLearn/day06/04myLogger"

// 测试自建日志库
func main() {
	logger := myLogger.NewLog()
	logger.Debug("这是一条Debug日志")
	logger.Info("这是一条Info日志")
	logger.Warning("这是一条Warning日志")
	logger.Error("这是一条Error日志")
	logger.Fatal("这是一条Fatal日志")

}
