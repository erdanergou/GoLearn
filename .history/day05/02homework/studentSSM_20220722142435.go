package main

import "fmt"

type student struct {
	id   int
	name string
}

type studentSSM struct {
	allstudent map[int]*student
}

// 查看学生
func (sm studentSSM) selectAllStudent() {
	for key, value := range sm.allstudent {
		fmt.Printf("学号：%d, 姓名：%s\n")key, value)
	}
}

// 增加学生
func (sm studentSSM) insertStudent() {
	fmt.Println("place enter new student message")
	var stu student
	fmt.Scanf("%d %s", &stu.id, &stu.name)
	sm.allstudent[stu.id] = &stu
}

// 修改学生
func (sm studentSSM) updateStudent() {

}

// 删除学生
func (sm studentSSM) deleteStudent() {

}
