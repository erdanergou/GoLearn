package table

type Company_Job struct {
	cid int 
}

func (cj *Company_Job) TableName() string {
	return "company_job"
}
