package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/shima004/slackclone/model"
	"github.com/shima004/slackclone/usecase"
)

type MessageHandler struct {
	MessageUseCase usecase.MessageUsecase
}

func NewMessageHandler(g *echo.Group, mu usecase.MessageUsecase) {
	handler := &MessageHandler{mu}
	g.GET("/messages", handler.FetchMessages)
	g.POST("/messages", handler.PostMessage)
	g.DELETE("/messages/:id", handler.DeleteMessage)
	g.PUT("/messages", handler.UpdateMessage)
}

func (mh *MessageHandler) FetchMessages(c echo.Context) error {
	sChannelID := c.QueryParam("channel_id")
	channelID, err := StringToUint(sChannelID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	sLimit := c.QueryParam("limit")
	limit, err := strconv.Atoi(sLimit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	sOffset := c.QueryParam("offset")
	offset, err := strconv.Atoi(sOffset)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	messages, err := mh.MessageUseCase.FetchMessages(ctx, channelID, limit, offset)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, messages)
}

func (mh *MessageHandler) PostMessage(c echo.Context) error {
	var message model.Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if ok, err := isRequestValid(&message); !ok {
		return c.JSON(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	err := mh.MessageUseCase.PostMessage(ctx, message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, message)
}

func (mh *MessageHandler) DeleteMessage(c echo.Context) error {
	sMessageID := c.Param("id")
	messageID, err := StringToUint(sMessageID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = mh.MessageUseCase.DeleteMessage(ctx, messageID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "Deleted")
}

func (mh *MessageHandler) UpdateMessage(c echo.Context) error {
	var message model.Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if ok, err := isRequestValid(&message); !ok {
		return c.JSON(http.StatusBadRequest, err)
	}

	if ok := isIncludeId(&message); !ok {
		return c.JSON(http.StatusBadRequest, "id is required")
	}

	ctx := c.Request().Context()
	err := mh.MessageUseCase.UpdateMessage(ctx, message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "Updated")
}

func isIncludeId(m *model.Message) bool {
	return m.ID != 0
}
