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
	name,ok := sm.allstudent[id]
	if ok{
		stu := sm[id]
		fmt.Scanf("place enter new name %s",)
	}
}

// 删除学生
func (sm studentSSM) deleteStudent() {

}
