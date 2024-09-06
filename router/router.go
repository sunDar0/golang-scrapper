package router

import (
	"github.com/labstack/echo/v4"

	"github.com/sunDar0/learngo/interfaces/controllers"
)

func SetupRouter(cmdController *controllers.CommandController, queryController *controllers.QueryController) *echo.Echo {
	e := echo.New()
	e.GET("/users", queryController.GetUsers)
	e.POST("/users", cmdController.CreateUser)
	e.GET("/users/:id", queryController.GetUser)
	return e
}
