package main

import (
	"fmt"
	"strings"
)

func main() {

	var path = "\"D:\\GoCode\\src\\GoLearn\\day01\\08string\""
	fmt.Printf("%s\n", path)

	s := "I'm Ok"
	fmt.Println(s)

	// 多行字符串 ``原样输出
	s2 := `卷帘西风
人比黄花瘦`

	fmt.Println(s2)

	//字符串相关操作
	fmt.Println(len(s))

	//字符串拼接

	firstName := "张"
	lastName := "三"
	ss := firstName + lastName
	fmt.Println(ss)

	ss1 := fmt.Sprintf("%s%s", firstName, lastName)

	fmt.Println(ss1)

	// 分割
	ret := strings.Split(path, "\\")
	fmt.Println(ret)

	// 包含
	fmt.Println(strings.Contains(s, "Ok"))

	// 前后缀判断
	// 前缀
	fmt.Println(strings.HasPrefix(s, "I'm"))
	// 后缀
	fmt.Println(strings.HasSuffix(s, "Ok"))

	s3 := "abcde"
	fmt.Println(strings.Index(s3,"c"))
	fmt.Println(strings.LastIndex(s3,"d"))

	// 拼接
	fmt.Println(strings.Join(ret,"+"))

}
