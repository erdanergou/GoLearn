package main

// 原子操作
// 针对整数数据类型（int32、uint32、int64、uint64）我们还可以使用原子操作来保证并发安全，通常直接使用原子操作比使用锁操作效率更高。

var x int64 = 0

func add() {
	x++
}

func main() {
	
}
