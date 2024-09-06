package command

import "context"

type JobCommandHandler struct {
	service *JobCommandService
}

func NewJobCommandHandler(service *JobCommandService) *JobCommandHandler {
	return &JobCommandHandler{service: service}
}

func (h *JobCommandHandler) HandleCreateJob(ctx context.Context, cmd CreateJobCommand) error {
	return h.service.CreateJob(ctx, cmd)
}
