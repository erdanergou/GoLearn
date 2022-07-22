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

	// var (
	// 	id   int
	// 	name string
	// )
	// fmt.Println("place enter new student id")
	// fmt.Scanln(&id)
	// fmt.Println("place enter new student name")
	// fmt.Scanln(&name)
	// stu := student{
	// 	id:   id,
	// 	name: name,
	// }

	var stu = new(student)
	fmt.Println("place enter new student id")
	fmt.Scanln(&stu.id)
	fmt.Println("place enter new student name")
	fmt.Scanln(&stu.name)
	sm.allstudent[stu.id] = stu
	fmt.Println("insert ok")
}

// 修改学生
func (sm studentSSM) updateStudent() {
	var id int
	fmt.Println("place enter update student id")
	fmt.Scanln(&id)
	stu, ok := sm.allstudent[id]
	if ok {
		fmt.Println("place enter new name")
		fmt.Scanln(&stu.name)
	} else {
		fmt.Printf("no student by id %d\n", id)
	}
}

// 删除学生
func (sm studentSSM) deleteStudent() {
	var id int
	fmt.Println("place enter delete student id")
	fmt.Scanln(&id)
	stu, ok := sm.allstudent[id]
	if ok {
		fmt.Printf("student %s will delete\n", stu.name)
		delete(sm.allstudent, id)
	} else {
		fmt.Printf("no student by id %d\n", id)
	}
}
