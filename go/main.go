package main

import (
	"github.com/labstack/echo/v4"
	"github.com/shima004/chat-server/controllers/web/handler"
	"github.com/shima004/chat-server/frameworks/web/router"
	"github.com/shima004/chat-server/gateways/datasource/mysqlimpl"
	sqlDB "github.com/shima004/chat-server/gateways/infra"
	"github.com/shima004/chat-server/gateways/repository/channel"
	"github.com/shima004/chat-server/gateways/repository/message"
	"github.com/shima004/chat-server/usecases/interactor"
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
	cds := mysqlimpl.NewMysqlChannel(db)
	mrepo := message.NewMessageRepo(mds)
	crepo := channel.NewChannelRepo(cds)
	mu := interactor.DefaultMessageUsecase{MessageRepository: mrepo, ChannelRepository: crepo, ContextTimeout: 10}
	cu := interactor.DefaultChannelUsecase{ChannelRepository: crepo, ContextTimeout: 10}
	mh := handler.NewMessageHandler(&mu)
	ch := handler.NewChannelHandler(&cu)
	router.NewChannelHandler(g, ch)
	router.NewMessageHandler(g, mh)
	e.Logger.Fatal(e.Start(":8080"))
}
