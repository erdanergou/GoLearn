package table

type Company_Job struct {
	Cid int `gorm:"cid"`
}

func (cj *Company_Job) TableName() string {
	return "company_job"
}
