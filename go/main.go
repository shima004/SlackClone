package main

import (
	"github.com/labstack/echo/v4"
	"github.com/shima004/slackclone/delivery"
	repo "github.com/shima004/slackclone/repository/mysql"
	"github.com/shima004/slackclone/usecase"
)

func main() {
	db, sqlDB, err := repo.Connect()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	e := echo.New()
	g := e.Group("/api")
	repo := repo.NewMysqlMessageRepository(db)
	mu := usecase.DefaultMessageUsecase{MessageRepository: repo}
	delivery.NewMessageHandler(g, &mu)

	e.Logger.Fatal(e.Start(":8080"))
}
