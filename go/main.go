package main

import (
	"github.com/labstack/echo/v4"
	"github.com/shima004/slackclone/controllers/web/handler"
	"github.com/shima004/slackclone/frameworks/web/router"
	"github.com/shima004/slackclone/gateways/datasource/mysqlimpl"
	sqlDB "github.com/shima004/slackclone/gateways/infra"
	"github.com/shima004/slackclone/gateways/repository/channel"
	"github.com/shima004/slackclone/gateways/repository/message"
	"github.com/shima004/slackclone/usecases/interactor"
)

func main() {
	db, sqlDB, err := sqlDB.Connect()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	e := echo.New()
	g := e.Group("/api")
	mds := mysqlimpl.NewMysqlMessage(db)
	mrepo := message.NewMessageRepo(mds)
	cds := mysqlimpl.NewMysqlChannel(db)
	crepo := channel.NewChannelRepo(cds)
	mu := interactor.DefaultMessageUsecase{MessageRepository: mrepo, ChannelRepository: crepo, ContextTimeout: 10}
	mh := handler.NewMessageHandler(&mu)
	router.NewMessageHandler(g, mh)
	e.Logger.Fatal(e.Start(":8080"))
}
