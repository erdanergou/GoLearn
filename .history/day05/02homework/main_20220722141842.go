package main

import (
	"fmt"
	"os"
)

// 结构体版学生管理系统

// 菜单函数

func showMenu() {

	fmt.Println("welcome ssm!")
	fmt.Println(`1.showAll   2.updateStudent  3.insertStudent   4.deleteStudent   5.exit`)

}


var ssm = studentSSM{
	
}
func main() {
	
	for {
		var choose int
		showMenu()
		fmt.Println("place enter you chose!")
		fmt.Scanln(&choose)
		switch choose {
		case 5:
			os.Exit(1)
		}
	}
}
