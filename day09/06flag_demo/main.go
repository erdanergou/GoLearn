package main

import (
	"flag"
	"fmt"
)

// flag 获取命令行参数
func main() {
	// 创建一个标志位
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 19, "年龄")
	married := flag.Bool("married", false, "婚姻状况")

	// var name string
	// flag.StringVar(&name, "name", "张三", "姓名")
	// 使用flag
	flag.Parse() // 先解析在使用
	// fmt.Println(name)
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*married)

	fmt.Println(flag.Args())  //返回命令行参数后的其他参数，以[]string类型
	fmt.Println(flag.NArg())  // 返回命令行参数后的其他参数类型
	fmt.Println(flag.NFlag()) //返回使用的命令行参数
}
