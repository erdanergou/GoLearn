package table

//  Person 用户信息
type Person struct {
	Name string `gorm:"name" json:"name"`
	Age  int    `gorm:"age" json:"age"`
	Id   int    `gorm:"id" json:"id"`
	Jobs []Job  `gorm:"job" json:"job"`
}

func (p *Person) TableName() string {
	return "person"
}
