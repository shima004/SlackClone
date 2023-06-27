package handler

import "github.com/labstack/echo/v4"

type IMessageHandler interface {
	FetchMessages(c echo.Context) error
	PostMessage(c echo.Context) error
	DeleteMessage(c echo.Context) error
	UpdateMessage(c echo.Context) error
}
