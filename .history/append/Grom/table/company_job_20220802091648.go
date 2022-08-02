package table

type Company_Job struct {
	cid 
}

func (cj *Company_Job) TableName() string {
	return "company_job"
}
