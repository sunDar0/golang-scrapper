package query

import "context"

type UserQueryHandler struct {
	service *UserQueryService
}

func NewUserQueryHandler(service *UserQueryService) *UserQueryHandler {
	return &UserQueryHandler{service: service}
}

func (h *UserQueryHandler) HandleGetUser(ctx context.Context, query GetUserQuery) (*UserDto, error) {
	return h.service.GetUser(ctx, query)
}
func (h *UserQueryHandler) HandleGetUsers(ctx context.Context) ([]*UserDto, error) {
	return h.service.GetUsers(ctx)
}
