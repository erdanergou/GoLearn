package mytable

type Company_Job struct {
	Cid int `gorm:"Cid"`
	Jid int `gorm:"job_id"`
}

func (cj *Company_Job) TableName() string {
	return "company_job"
}
