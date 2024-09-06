package models

// User 도메인 엔티티
type Job struct {
	Id        string
	Company   string
	Title     string
	WorkPlace string
	Career    string
	Summary   []string
}

func NewJob(id, company, title, workPlace, career string, summary []string) (*Job, error) {

	return &Job{
		Id:        id,
		Company:   company,
		Title:     title,
		WorkPlace: workPlace,
		Career:    career,
		Summary:   summary,
	}, nil
}
