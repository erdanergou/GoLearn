package main

import "fmt"

/*
包的定义 --> package关键字，包名通常与目录名一致,不能用-
包的导入 --> import关键字
	1.单行导入
	2.多行导入
	3.给导入的包起别名
	4.匿名导入
	5.包导入路径是从&GOPATH/src后面的路径开始写起
	6.包不支持循环导入
包中的标识符可见性-->标识符首字母大写表示对外可见

init()
	1.包导入时会自动执行
	2.一个包里面只有init()
	3.init()没有参数也没有返回值也不能调用它
	4.多个包的init执行顺序
	5.一把用于做一些初始化操作


接口
	接口是一种类型，一种抽象类型
	定义：
		type 接口名 interface{
			方法名（参数）返回值
		}
	实现了接口的所有方法就实现了这个接口
	实现了接口就可以当成这个接口类型的变量
接口变量
	实现了一个万能的变量，可以保存所有实现了这个接口的类型的值
	通常作为函数的参数出现
空接口
	interface{}
	接口中没有定义任何方法，也就是说任意类型都实现了空接口-->任何类型都可以存到空接口变量中
	作为函数参数-->fmt.Println()
	map[string]interface{}
接口底层
	动态类型
	动态值
类型断言
	类型断言的前提是一定要是一个接口类型的变量
	x.(T)
	使用switch来做类型断言

*/

//  类型断言
func main() {
	var a interface{} //定义一个空接口变量
	a = 100
	// 如何判断a保存的值的具体类型
	// 类型断言
	// 1.x.(T)
	v, ok := a.(int8)
	if !ok {
		fmt.Println("猜对了,a是int8", v)
		return
	} else {
		fmt.Println("猜错了,a不是int8", v)
	}

	//switch
	switch v2 := a.(type) {
	case int8:
		fmt.Println("int8", v2)
	case int16:
		fmt.Println("int16", v2)
	case int:
		fmt.Println("int", v2)
	case string:
		fmt.Println("string", v2)
	default:
		fmt.Println("滚~")
	}
}

/*
文件操作

读文件：
	file,err := os.OpenFile(path)
	file.Read()
	bufio
	ioutil


写文件:
	file,err := os.OpenFile(path)
	file.Write()/file.WriteString()

	buifo
	ioutil


	在读写文件时，需要在最后对文件进行关闭，此时可以使用defer关键字
	此时要注意：因为在OpenFile的时候获取到了两个返回值file，以及err，同时会进行判断是否有err发生，此时defer file.Close要放在判断err的最后
	一是：因为在判断是如果有错误会直接return，此时不会执行下面的语句
	二是：因为如果将defer file.Close放在前面，若真的出现错误，此时的file为nil，无法调用file.Close()
	二附加：当代码执行到defer file.Close() 的时候会将file进行赋值，此时若为nil，在执行的最后为nil.Close()
*/
