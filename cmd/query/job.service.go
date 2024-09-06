package query

import (
	"context"

	domain "github.com/sunDar0/learngo/domain/repositories"
)

type JobQueryService struct {
	repository domain.JobRepository
}

func NewJobQueryService(repo domain.JobRepository) *JobQueryService {
	return &JobQueryService{repository: repo}
}

func (s *JobQueryService) GetJob(ctx context.Context, query GetJobQuery) (*JobDto, error) {
	job, err := s.repository.FindByID(ctx, query.JobId)
	if err != nil {
		return nil, err
	}

	return &JobDto{Id: job.Id}, nil
}

func (s *JobQueryService) GetJobs(ctx context.Context) ([]*JobDto, error) {
	jobs, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var jobDto []*JobDto
	for _, job := range jobs {
		jobDto = append(jobDto, &JobDto{Id: job.Id})
	}
	return jobDto, nil
}
