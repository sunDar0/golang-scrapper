package command

import (
	"context"
	"fmt"

	"github.com/sunDar0/learngo/domain/models"
	domain "github.com/sunDar0/learngo/domain/repositories"
)

type UserCommandService struct {
	repository domain.UserRepository
}

func NewUserCommandService(repo domain.UserRepository) *UserCommandService {
	return &UserCommandService{repository: repo}
}

func (s *UserCommandService) CreateUser(ctx context.Context, cmd CreateUserCommand) error {
	// 비즈니스 로직 처리
	fmt.Println(cmd)
	user, _ := models.NewUser(cmd.Name, cmd.Email)
	return s.repository.Save(ctx, user)
}
