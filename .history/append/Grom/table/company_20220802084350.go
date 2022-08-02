package table

type Company struct {
	Id      int    `gorm:"id"`
	Name    string `gorm:"name"`
	Persons []Person `gorm:"for`
}

func (c *Company) TableName() string {
	return "company"
}
