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
		fmt.Println(key, value)
	}
}

// 增加学生
func (sm studentSSM) insertStudent() {
	
}

// 修改学生
func (sm studentSSM) updateStudent() {

}

// 删除学生
func (sm studentSSM) deleteStudent() {

}
