package table

type Company struct {
	Id      int    `gorm:"id"`
	Name    string `gorm:"name"`
	Persons []Person `gorm:"foreignkey:Cate
}

func (c *Company) TableName() string {
	return "company"
}
