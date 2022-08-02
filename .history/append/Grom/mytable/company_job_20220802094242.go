package mytable

type Company_Job struct {
	Cid int `gorm:"company_id"`
	Jid int `gorm:"job_id"`
}

func (cj *Company_Job) TableName() string {
	return "company_job"
}
