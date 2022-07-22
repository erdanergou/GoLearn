package calc

// 表中的标识符（变量名、函数名、结构体、接口等）如果是小写，表示私有（只能在当前包中使用）
// 首字母大写的标识符可以
func add(x, y int) int {
	return x + y
}
