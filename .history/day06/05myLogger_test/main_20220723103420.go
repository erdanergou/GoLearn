package main

import "GoLearn/day0"

// 测试自建日志库
func main() {
	logger := myLogger.NewLog()
	logger.Debug("这是一条Debug日志")
}
