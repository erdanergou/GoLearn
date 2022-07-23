package main

/*
time  
时间类型 :
	time.time:time.Now()当前时间
	时间戳:time.Now().Unix:1970.1.1到现在的秒数
	纳秒数:time.Now().UnixNano():1970.1.1到现在的纳秒数
时间间隔类型:
	time.Duration:
时间操作:
	时间对象 +-*\ 时间对象
	after/before
时间格式化:2006-01-02 03:04:05 || 2006-01-02 15:04:05 000
定时器:
	time.Tick(时间)

解析字符串格式的时间(时区) day06/02time 里的f1

反射:
接口类型的变量底层是分为动态类型和动态值
反射的应用:json解析

*/

type student struct{
	Name string `json:"name"`
	Age int `json:"age"`
}


func main() {
	str := `{"name":"张三","age}`
}
