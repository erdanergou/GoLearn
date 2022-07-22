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
	for _, value := range sm.allstudent {
		fmt.Printf("学号：%d, 姓名：%s\n", value.id, value.name)
	}
}

// 增加学生
func (sm studentSSM) insertStudent() {
	fmt.Println("place enter new student message")

	var (
		id   int
		name string
	)
	fmt.Println("place enter new student id")
	fmt.Scanln(&id)
	fmt.Println("place enter new student name")
	fmt.Scanln(&name)
	stu := student{
		id:   id,
		name: name,
	}
	sm.allstudent[stu.id] = &stu
}

// 修改学生
func (sm studentSSM) updateStudent() {
	fmt.Println("place enter update student id")
	var id int
	fmt.Scanln(&id)

}

// 删除学生
func (sm studentSSM) deleteStudent() {

}
