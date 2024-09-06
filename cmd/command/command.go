package command

// CreateUserCommand 는 사용자 생성을 위한 커맨드입니다.
type CreateUserCommand struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateJobCommand struct {
	Id        string
	Company   string
	Title     string
	WorkPlace string
	Career    string
	Summary   []string
}
