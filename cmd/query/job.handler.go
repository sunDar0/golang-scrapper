package query

import "context"

type JobQueryHandler struct {
	service *JobQueryService
}

func NewJobQueryHandler(service *JobQueryService) *JobQueryHandler {
	return &JobQueryHandler{service: service}
}

func (h *JobQueryHandler) HandleGetJob(ctx context.Context, query GetJobQuery) (*JobDto, error) {
	return h.service.GetJob(ctx, query)
}
func (h *JobQueryHandler) HandleGetJobs(ctx context.Context) ([]*JobDto, error) {
	return h.service.GetJobs(ctx)
}
