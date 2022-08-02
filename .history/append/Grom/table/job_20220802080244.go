package table

import "github.com/jinzhu/gorm"

//  Person 用户信息
type Job struct {
	gorm.Model
	Name   string `gorm:"name" json:"name"`
	Salary int    `gorm:"salary" json:"salary"`
	Id     int    `gorm:"id" json:"id"`
}

func (j *Job) TableName() string {
	return "job"
}
