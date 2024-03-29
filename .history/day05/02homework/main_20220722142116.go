package main

import (
	"fmt"
	"os"
)

// 结构体版学生管理系统
var ssm studentSSM // 声明全局变量
// 菜单函数

func showMenu() {

	fmt.Println("welcome ssm!")
	fmt.Println(`1.showAll   2.updateStudent  3.insertStudent   4.deleteStudent   5.exit`)

}

func main() {

	for {
		ssm = studentSSM{
			allstudent: make(map[int]*student, 32),
		}
		var choose int
		showMenu()
		fmt.Println("place enter you chose!")
		fmt.Scanln(&choose)
		switch choose {
		case 5:
			os.Exit(1)
		case 1:
			ssm.selectAllStudent()
		case 2:
			ssm.updateStudent()
		case 3:
			ssm.insertStudent()
		case 4:
		}
	}
}
