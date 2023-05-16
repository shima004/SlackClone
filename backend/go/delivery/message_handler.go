package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shima004/slackclone/usecase"
)

type MessageHandler struct {
	MessageUseCase usecase.MessageUsercase
}

func NewMessageHandler(e *echo.Echo, mu usecase.MessageUsercase) {
	handler := &MessageHandler{mu}
	e.GET("/api/messages", handler.FetchMessages)
}

func (mh *MessageHandler) FetchMessages(c echo.Context) error {
	auther := c.QueryParam("auther")
	ctx := c.Request().Context()
	messages, err := mh.MessageUseCase.FetchMessages(ctx, auther)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(http.StatusOK, messages)
}
