package query

import (
	"context"

	domain "github.com/sunDar0/learngo/domain/repositories"
)

type UserQueryService struct {
	repository domain.UserRepository
}

func NewUserQueryService(repo domain.UserRepository) *UserQueryService {
	return &UserQueryService{repository: repo}
}

func (s *UserQueryService) GetUser(ctx context.Context, query GetUserQuery) (*UserDto, error) {
	user, err := s.repository.FindByID(ctx, query.UserID)
	if err != nil {
		return nil, err
	}

	return &UserDto{Name: user.Name, Email: user.Email}, nil
}

func (s *UserQueryService) GetUsers(ctx context.Context) ([]*UserDto, error) {
	users, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var userDto []*UserDto
	for _, user := range users {
		userDto = append(userDto, &UserDto{ID: user.ID, Name: user.Name, Email: user.Email})
	}
	return userDto, nil
}
