package router

import (
	"github.com/labstack/echo/v4"
	"github.com/shima004/slackclone/frameworks/web/handler"
)

type MessageRouter struct {
	MessageHandler handler.IMessageHandler
}

func NewMessageHandler(g *echo.Group, mh handler.IMessageHandler) {
	router := &MessageRouter{mh}
	g.GET("/messages", router.MessageHandler.FetchMessages)
	g.POST("/messages", router.MessageHandler.PostMessage)
	g.DELETE("/messages/:id", router.MessageHandler.DeleteMessage)
	g.PUT("/messages", router.MessageHandler.UpdateMessage)
}
