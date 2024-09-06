package command

import "context"

type UserCommandHandler struct {
	service *UserCommandService
}

func NewUserCommandHandler(service *UserCommandService) *UserCommandHandler {
	return &UserCommandHandler{service: service}
}

func (h *UserCommandHandler) HandleCreateUser(ctx context.Context, cmd CreateUserCommand) error {
	return h.service.CreateUser(ctx, cmd)
}
