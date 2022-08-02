package table

type Company struct {
	Id      int      `gorm:"id"`
	Name    string   `gorm:"name"`
	Persons []Person `gorm:"foreignkey:Cid association_foreignkey:Name`
}

func (c *Company) TableName() string {
	return "company"
}
