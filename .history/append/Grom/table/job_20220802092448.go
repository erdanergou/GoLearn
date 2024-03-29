package mytable

//  Person 用户信息
type Job struct {
	Name     string `gorm:"name" json:"name"`
	Salary   int    `gorm:"salary" json:"salary"`
	Id       int    `gorm:"id" json:"id"`
	Companys []Company `gorm:"many2many:company_job"`
}

func (j *Job) TableName() string {
	return "job"
}
