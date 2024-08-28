// Package main is the main package of the program.
package main

import (
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sunDar0/learngo/common"
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
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal((e.Start((":1323"))))
}
