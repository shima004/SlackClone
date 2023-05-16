package main

import (
	"github.com/labstack/echo/v4"
	"github.com/shima004/slackclone/controller"
	"github.com/shima004/slackclone/repository"
)

func main() {
	e := echo.New()
	g := e.Group("/api")

	controller := controller.Controller{Repository: &repository.DefaultRepository{}}
	g.GET("/messages", controller.GetAllMessages)

	e.Logger.Fatal(e.Start(":8080"))
}
