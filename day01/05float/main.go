package main

import (
	"fmt"
	"strconv"
)

// 浮点数

func main() {
	// math.MaxFloat32 // float32的最大值
	f1 := 1.23456
	fmt.Printf("%T\n", f1) // 默认Go语言中的小数都是float64类型
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2) // 显式声明float32类型
	// f1 = float64(f2) //float32类型的值不能直接赋值给float64

	//类型转换
	// 1.string转整形
	s := "12"
	disInt, _ := strconv.Atoi(s)
	fmt.Printf("disInt %d\n", disInt) // 将字符串13转换为int类型的数值13
	// 参数1：带转换字符串，
	// 参数2：基于几进制，值可以是0,8,16,32,64
	// 参数3：要转成哪个int类型：可以是0、8、16、32、64，分别对应 int,int8,int16,int32,int64
	disInt64, _ := strconv.ParseInt(s, 10, 64)
	fmt.Printf("disInt64 %d\n", disInt64)

	// ParseFloat 将字符串转换为浮点数
	// str：要转换的字符串
	// bitSize：指定浮点类型（32:float32、64:float64）
	// 如果 str 是合法的格式，而且接近一个浮点值，
	// 则返回浮点数的四舍五入值（依据 IEEE754 的四舍五入标准）
	// 如果 str 不是合法的格式，则返回“语法错误”
	// 如果转换结果超出 bitSize 范围，则返回“超出范围”
	//到float64
	distFloat64, _ := strconv.ParseFloat(s, 64)
	fmt.Printf("disInt64 %f\n", distFloat64)
	//到float32
	distFloat32, _ := strconv.ParseFloat(s, 32)
	fmt.Printf("disInt64 %f\n", distFloat32)

	// 整形转string
	var num int64 = 12
	distStr := strconv.FormatInt(num, 10)
	fmt.Printf("%s\n", distStr)

	// float转string
	// FormatFloat 将浮点数 f 转换为字符串值
	// f：要转换的浮点数
	// fmt：格式标记（b、e、E、f、g、G）
	// prec：精度（数字部分的长度，不包括指数部分）
	// bitSize：指定浮点类型（32:float32、64:float64）
	// 格式标记：
	// 'b' (-ddddp±ddd，二进制指数)
	// 'e' (-d.dddde±dd，十进制指数)
	// 'E' (-d.ddddE±dd，十进制指数)
	// 'f' (-ddd.dddd，没有指数)
	// 'g' ('e':大指数，'f':其它情况)
	// 'G' ('E':大指数，'f':其它情况)
	// 如果格式标记为 'e'，'E'和'f'，则 prec 表示小数点后的数字位数
	// 如果格式标记为 'g'，'G'，则 prec 表示总的数字位数（整数部分+小数部分）

	str1 := strconv.FormatFloat(11.34, 'E', -1, 32)
	str2 := strconv.FormatFloat(10.55, 'E', -1, 64)
	fmt.Println(str1, str2) //1.134E+01  1.055E+01

	//解析转换后的string变量str为float
	h, _ := strconv.ParseFloat(str1, 32)
	fmt.Println(h) //11.34000015258789
	h, _ = strconv.ParseFloat(str2, 64)
	fmt.Println(h) //10.55

	str := strconv.FormatFloat(1.1, 'f', 0, 64)
	fmt.Println(str) // 1

	str = strconv.FormatFloat(1.1, 'f', 1, 64)
	fmt.Println(str) // 1.1

	str = strconv.FormatFloat(1.1, 'f', 2, 64)
	fmt.Println(str) // 1.10

	str = strconv.FormatFloat(1.1, 'f', -1, 64)
	fmt.Println(str) // 1.1

	
}
