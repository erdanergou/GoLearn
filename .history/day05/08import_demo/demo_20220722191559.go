package main  // 只有main包才能编译成可执行文件

import ("github.com/GoLearn/day")

fun main(){
	ret := calc.Add(1,2)
	fmt.Println(ret)
}