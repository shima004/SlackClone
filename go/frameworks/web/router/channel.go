package router

import (
	"github.com/labstack/echo/v4"
	"github.com/shima004/slackclone/frameworks/web/handler"
)

type ChannelRouter struct {
	ChannelHandler handler.IChannelHandler
}

func NewChannelHandler(g *echo.Group, ch handler.IChannelHandler) {
	router := &ChannelRouter{ch}
	g.POST("/channels", router.ChannelHandler.PostChannel)
	g.DELETE("/channels/:id", router.ChannelHandler.DeleteChannel)
}
