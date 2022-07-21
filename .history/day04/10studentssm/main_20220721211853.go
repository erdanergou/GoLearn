package main

import (
	"fmt"
	"os"
)

/*函数版学生管理系统
写一个系统可以查看\新增\删除学生
*/
type student struct {
	id   int
	name string
}

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

var (
	allstudent map[int64](*student) //变量声明
)

func main() {
	// 1.打印菜单
	allstudent = make(map[int64]*student, 32) //初始化,开辟内存空间
	fmt.Println("欢迎使用学生管理系统")
	for {
		fmt.Println(`
	1.查看所有学生  2.新增学生  3.删除学生  4.退出
请输入要进行的操作编号:
	`)
		// 2.等待用户选择
		var useId int

		fmt.Scanln(&useId)
		switch useId {
		case 1:
			selectSudent()
		case 2:
			insertSudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("请检查输入是否正确")
		}
	}
	// 3.执行对应函数
}

func selectSudent() {
	for key, value := range allstudent {
		fmt.Printf("学号:%d 姓名:%s", key, value.name)
	}
}

func insertSudent() {
	var name string
	var id int64
	fmt.Scanf("%s %d", &name, &id)
	stu := newStudent(id, name)
	allstudent[id] = stu
}

func deleteStudent() {
	fmt.Println("删除学生")
}
