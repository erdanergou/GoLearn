package mytable

type Company_Job struct {
	Cid int `gorm:"cid"`
	Jid int `gorm:"jid"`
}

func (cj *Company_Job) TableName() string {
	return "company_job"
}
