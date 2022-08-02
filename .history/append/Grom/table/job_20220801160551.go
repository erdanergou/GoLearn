package table

//  Person 用户信息
type Person struct {
	Name string `gorm:"name" json:"name"`
	Age  int    `gorm:"age" json:"age"`
	Id   int    `gorm:"id" json:"id"`
	Work int    `gorm:"work" json:"work"`
}

func (p *Person) TableName() string {
	return "person"
}