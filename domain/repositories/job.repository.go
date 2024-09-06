package domain

import (
	"context"

	"github.com/sunDar0/learngo/domain/models"
)

// JobRepository 인터페이스 정의
type JobRepository interface {
	Save(ctx context.Context, job *models.Job) error
	FindByID(ctx context.Context, id string) (*models.Job, error)
	FindAll(ctx context.Context) (map[string]*models.Job, error)
}
