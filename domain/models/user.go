package models

import (
	"errors"

	"github.com/google/uuid"
)

// User 도메인 엔티티
type User struct {
	ID    string
	Name  string
	Email string
}

// NewUser는 새로운 User 객체를 생성합니다.
func NewUser(name, email string) (*User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email cannot be empty")
	}

	return &User{
		ID:    uuid.New().String(), // UUID를 사용하여 고유한 ID 생성
		Name:  name,
		Email: email,
	}, nil
}

// UpdateEmail은 사용자의 이메일을 변경하는 비즈니스 로직입니다.
func (u *User) UpdateEmail(newEmail string) error {
	if newEmail == "" {
		return errors.New("email cannot be empty")
	}
	u.Email = newEmail
	return nil
}
