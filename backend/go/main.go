package main

import (
	"github.com/labstack/echo/v4"
	"github.com/shima004/slackclone/delivery"
	"github.com/shima004/slackclone/repository"
	"github.com/shima004/slackclone/usecase"
)

func main() {
	e := echo.New()
	g := e.Group("/api")

	repo := repository.DefaultMessageRepository{}
	mu := usecase.DefaultMessageUsercase{MessageRepository: &repo}
	delivery.NewMessageHandler(g, &mu)

	e.Logger.Fatal(e.Start(":8080"))
}
