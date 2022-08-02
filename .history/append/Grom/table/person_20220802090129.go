package table

//  Person 用户信息
type Person struct {
	Name  string `gorm:"name" json:"name"`
	Age   int    `gorm:"age" json:"age"`
	Id    int    `gorm:"id" json:"id"`
	Jobid int    `gorm:"jobid" `
	Job   Job    `gorm:"foreignKey:jobid;AssociationForeignKey:id"`
	Cid   int    `gorm:"cid"`
}

func (p *Person) TableName() string {
	return "person"
}
