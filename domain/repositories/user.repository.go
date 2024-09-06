package domain

import (
	"context"

	"github.com/sunDar0/learngo/domain/models"
)

// UserRepository 인터페이스 정의
type UserRepository interface {
	Save(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id string) (*models.User, error)
	FindAll(ctx context.Context) (map[string]*models.User, error)
}
