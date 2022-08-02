package table

type Company struct {
	Id      int      `gorm:"id"`
	Name    string   `gorm:"name"`
	Persons []Person `gorm:"foreignkey:Cid;association_foreignkey:Id`
}

func (c *Company) TableName() string {
	return "company"
}
