package command

import (
	"context"
	"fmt"

	"github.com/sunDar0/learngo/domain/models"
	domain "github.com/sunDar0/learngo/domain/repositories"
)

type JobCommandService struct {
	repository domain.JobRepository
}

func NewJobCommandService(repo domain.JobRepository) *JobCommandService {
	return &JobCommandService{repository: repo}
}

func (s *JobCommandService) CreateJob(ctx context.Context, cmd CreateJobCommand) error {
	// 비즈니스 로직 처리
	fmt.Println(cmd)

	user, _ := models.NewJob(cmd.Id, cmd.Company, cmd.Title, cmd.WorkPlace, cmd.Career, cmd.Summary)
	return s.repository.Save(ctx, user)
}
