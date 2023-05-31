package main

import (
	"time"

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
	mRepo := repo.NewMysqlMessageRepository(db)
	cRepo := repo.NewMysqlChannelRepository(db)
	mu := usecase.DefaultMessageUsecase{MessageRepository: mRepo, ChannelRepository: cRepo, ContextTimeout: 10 * time.Second}
	cu := usecase.DefaultChannelUsecase{ChannelRepository: cRepo, ContextTimeout: 10 * time.Second}
	delivery.NewMessageHandler(g, &mu)
	delivery.NewChannelHandler(g, &cu)

	e.Logger.Fatal(e.Start(":8080"))
}
