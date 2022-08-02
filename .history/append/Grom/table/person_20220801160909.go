package table

import "github.com/casbin/casbin/v2/model"

//  Person 用户信息
type Person struct {
	Name string `gorm:"name" json:"name"`
	Age  int    `gorm:"age" json:"age"`
	Id   int    `gorm:"id" json:"id"`
	Jobs  model.job    `gorm:"job" json:"job"`
}

func (p *Person) TableName() string {
	return "person"
}
