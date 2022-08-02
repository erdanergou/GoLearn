package table

//  Person 用户信息
type Job struct {
	Name string `gorm:"name" json:"name"`
	Age  int    `gorm:"age" json:"age"`
	Id   int    `gorm:"id" json:"id"`
}

func (p *Person) TableName() string {
	return "person"
}