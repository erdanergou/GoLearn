package main

// 常量(定义后不能修改，程序运行期间不会改变的量)
const pi = 3.1415926 

// 批量声明常量
const (
	statusOk = 200
	notFound = 404
)

// 批量声明常量时，某行声明后没有赋值，默认与上行同值
const(
	n1 = 100
	n2
	n3 
)

//iota是go语言的常量计数器,在常量出现的时候置为0
const(
	a1 = iota //0
	a2
	a3
)


// 定义数量级
const(
	_ = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main(){
	println(n1)
	println(n2)
	println(n3)


	println(a1)
	println(a2)
	println(a3)
}