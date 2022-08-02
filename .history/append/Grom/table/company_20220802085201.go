package table

type Company struct {
	Id      int    `gorm:"id"`
	Name    string `gorm:"name"`
	Persons []Person `gorm:"ForeignKey:cid;AssociationForeignKey:id"`
}

func (c *Company) TableName() string {
	return "company"
}
