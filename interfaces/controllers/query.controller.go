package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sunDar0/learngo/cmd/query"
)

type QueryController struct {
	handler *query.UserQueryHandler
}

func NewQueryController(handler *query.UserQueryHandler) *QueryController {
	return &QueryController{handler: handler}
}

func (c *QueryController) GetUser(ctx echo.Context) error {
	userID := ctx.Param("id")
	result, err := c.handler.HandleGetUser(ctx.Request().Context(), query.GetUserQuery{UserID: userID})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, result)
}

func (c *QueryController) GetUsers(ctx echo.Context) error {
	result, err := c.handler.HandleGetUsers(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, result)
}
