package persistence

import (
	"context"
	"errors"
	"sync"

	"github.com/sunDar0/learngo/domain/models"
)

type InMemJobRepo struct {
	mu   sync.RWMutex
	jobs map[string]*models.Job
}

// NewUserRepository 는 새로운 InMemoryUserRepository 인스턴스를 반환합니다.
func NewJobRepository() *InMemJobRepo {
	return &InMemJobRepo{
		jobs: make(map[string]*models.Job),
	}
}

// Save 메서드는 사용자 데이터를 저장합니다.
func (r *InMemJobRepo) Save(ctx context.Context, job *models.Job) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.jobs[job.Id] = job
	return nil
}

// FindByID 메서드는 ID를 기반으로 사용자 데이터를 조회합니다.
func (r *InMemJobRepo) FindByID(ctx context.Context, id string) (*models.Job, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if job, exists := r.jobs[id]; exists {
		return job, nil
	}
	return nil, errors.New("job not found")
}

func (r *InMemJobRepo) FindAll(ctx context.Context) (map[string]*models.Job, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if len(r.jobs) > 0 {
		return r.jobs, nil
	}
	return nil, errors.New("users not found")
}
