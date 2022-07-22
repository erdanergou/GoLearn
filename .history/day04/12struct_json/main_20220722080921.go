package main

import (
	"encoding/json"
	"fmt"
)

//结构体与json
// 1. 序列化：把go语言中的结构体变量 --> json格式的字符串
// 2. 反序列化：json格式的字符串 --> go语言中能够识别的结构体变量

type person struct { //结构体字段的可见性，因为格式化的功能是json包进行的，其实是在json包里的marshal方法里把传入的结构体变量的所有东西都拿出来转成字符串
	Name string `json:"name",db:"name" "ini:name"` // 当使用json解析时，按照"name"去代替"Name"
	Age  int    //如果是小写，在当前包可以正常访问，但json包不可以
}

func main() {
	p1 := person{
		Name: "张三",
		Age:  17,
	}
	fmt.Println(p1)
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Printf("%#v", string(b))
}
