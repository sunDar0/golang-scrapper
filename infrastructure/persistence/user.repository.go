package persistence

import (
	"context"
	"errors"
	"sync"

	"github.com/sunDar0/learngo/domain/models"
)

type InMemoryUserRepository struct {
	mu    sync.RWMutex
	users map[string]*models.User
}

// NewUserRepository 는 새로운 InMemoryUserRepository 인스턴스를 반환합니다.
func NewUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*models.User),
	}
}

// Save 메서드는 사용자 데이터를 저장합니다.
func (r *InMemoryUserRepository) Save(ctx context.Context, user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.ID] = user
	return nil
}

// FindByID 메서드는 ID를 기반으로 사용자 데이터를 조회합니다.
func (r *InMemoryUserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if user, exists := r.users[id]; exists {
		return user, nil
	}
	return nil, errors.New("user not found")
}

func (r *InMemoryUserRepository) FindAll(ctx context.Context) (map[string]*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if len(r.users) > 0 {
		return r.users, nil
	}
	return nil, errors.New("users not found")
}
