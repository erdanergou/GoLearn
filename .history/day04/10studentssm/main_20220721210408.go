package main

import (
	"fmt"
	"os/user"
)

/*函数版学生管理系统
写一个系统可以查看\新增\删除学生
*/
type student struct {
	id   int
	name string
}

func main() {
	// 1.打印菜单
	fmt.Println("欢迎使用学生管理系统")
	fmt.Println(`
	1.查看所有学生  2.新增学生  3.删除学生
请输入要进行的操作编号:
	`)
	// 2.等待用户选择
	var useId int
	fmt.Scanln(useId)
	
	// 3.执行对应函数
}
