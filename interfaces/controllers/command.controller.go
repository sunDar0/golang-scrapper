package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sunDar0/learngo/cmd/command"
)

type CommandController struct {
	handler *command.UserCommandHandler
}

func NewCommandController(handler *command.UserCommandHandler) *CommandController {
	return &CommandController{handler: handler}
}

func (c *CommandController) CreateUser(ctx echo.Context) error {
	var cmd command.CreateUserCommand
	if err := ctx.Bind(&cmd); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	if err := c.handler.HandleCreateUser(ctx.Request().Context(), cmd); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusCreated)
}
