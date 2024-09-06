package query

// GetUserQuery 는 사용자 조회를 위한 쿼리입니다.
type GetUserQuery struct {
	UserID string `json:"id"`
}

// UserDTO 는 조회된 사용자 데이터를 반환하는 데이터 전송 객체입니다.
type UserDto struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GetUserQuery 는 사용자 조회를 위한 쿼리입니다.
type GetJobQuery struct {
	JobId string `json:"id"`
}

type JobDto struct {
	Id        string
	Company   string
	Title     string
	WorkPlace string
	Career    string
	Summary   []string
}
