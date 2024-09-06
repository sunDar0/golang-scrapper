// Package main is the main package of the program.
package main

import (
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sunDar0/learngo/cmd/command"
	"github.com/sunDar0/learngo/cmd/query"
	"github.com/sunDar0/learngo/common"
	"github.com/sunDar0/learngo/infrastructure/persistence"
	"github.com/sunDar0/learngo/interfaces/controllers"
	"github.com/sunDar0/learngo/router"
	"github.com/sunDar0/learngo/scrapper"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}
func handleScrape(c echo.Context) error {

	searchKeyword := c.FormValue("searchKeyword")
	searchKeyword = strings.ToLower(common.CleanString(searchKeyword))
	scrapper.Scrape(searchKeyword)

	defer os.Remove(searchKeyword + "_jobs.csv")
	return c.Attachment(searchKeyword+"_jobs.csv", searchKeyword+"_jobs.csv")
}

func main() {
	// e := echo.New()
	// e.GET("/", handleHome)
	// e.POST("/scrape", handleScrape)
	// e.Logger.Fatal((e.Start((":1323"))))

	userRepo := persistence.NewUserRepository()
	userCmdService := command.NewUserCommandService(userRepo)
	userQueryService := query.NewUserQueryService(userRepo)
	userCmdHandler := command.NewUserCommandHandler(userCmdService)
	userQueryHandler := query.NewUserQueryHandler(userQueryService)

	userCmdController := controllers.NewCommandController(userCmdHandler)
	userQueryController := controllers.NewQueryController(userQueryHandler)

	jobRepo := persistence.NewJobRepository()
	jobCmdService := command.NewJobCommandService(jobRepo)
	jobQueryService := query.NewJobQueryService(jobRepo)
	command.NewJobCommandHandler(jobCmdService)
	query.NewJobQueryHandler(jobQueryService)

	// 라우터 설정 및 서버 시작
	e := router.SetupRouter(userCmdController, userQueryController)
	e.Start(":8080")
}
