package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shima004/slackclone/entities"
	"github.com/shima004/slackclone/usecases/inputport"
	"github.com/shima004/slackclone/utility"
)

type MessageHandler struct {
	MessageInputPort inputport.MessageUsecase
}

func NewMessageHandler(mip inputport.MessageUsecase) *MessageHandler {
	return &MessageHandler{mip}
}

func (mh *MessageHandler) FetchMessages(c echo.Context) error {
	sChannelID := c.QueryParam("channel_id")
	sLimit := c.QueryParam("limit")
	sOffset := c.QueryParam("offset")

	channelID, err := utility.StringToUint(sChannelID)
	if err != nil {
		return err
	}
	limit, err := utility.StringToInt(sLimit)
	if err != nil {
		return err
	}
	offset, err := utility.StringToInt(sOffset)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	messages, err := mh.MessageInputPort.FetchMessages(ctx, channelID, limit, offset)
	if err != nil {
		return err
	}

	return c.JSON(200, messages)
}

func (mh *MessageHandler) PostMessage(c echo.Context) error {
	var message entities.Message
	if err := c.Bind(&message); err != nil {
		return err
	}
	if ok, err := utility.IsRequestValid(&message); !ok {
		return err
	}

	ctx := c.Request().Context()
	if err := mh.MessageInputPort.PostMessage(ctx, &message); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, message)
}

func (mh *MessageHandler) DeleteMessage(c echo.Context) error {
	sMessageID := c.Param("message_id")
	messageID, err := utility.StringToUint(sMessageID)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	if err := mh.MessageInputPort.DeleteMessage(ctx, messageID); err != nil {
		return err
	}
	return c.JSON(200, nil)
}

func (mh *MessageHandler) UpdateMessage(c echo.Context) error {
	var message entities.Message
	if err := c.Bind(&message); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if err := mh.MessageInputPort.UpdateMessage(ctx, &message); err != nil {
		return err
	}
	return c.JSON(200, message)
}
