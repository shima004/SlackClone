package handler

import (
	"github.com/labstack/echo/v4"
)

type IChannelHandler interface {
	PostChannel(c echo.Context) error
	DeleteChannel(c echo.Context) error
}
