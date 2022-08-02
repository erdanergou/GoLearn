package mytable

type Company struct {
	Id      int      `gorm:"id"`
	Name    string   `gorm:"name"`
	Persons []Person `gorm:"ForeignKey:cid"`
	Jobs    []Job    `gorm:"many2many:company_job"`
}

func (c *Company) TableName() string {
	return "company"
}
