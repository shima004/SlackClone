package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shima004/chat-server/entities"
	cerror "github.com/shima004/chat-server/entities/error"
	"github.com/shima004/chat-server/usecases/inputport"
	"github.com/shima004/chat-server/usecases/inputport/validation"
	"github.com/shima004/chat-server/utility"
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
		return c.JSON(http.StatusBadRequest, err)
	}
	limit, err := utility.StringToInt(sLimit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	offset, err := utility.StringToInt(sOffset)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	in := &validation.FatchMessagesInput{
		ChannelID: channelID,
		Limit:     limit,
		Offset:    offset,
	}

	ctx := c.Request().Context()
	messages, err := mh.MessageInputPort.FetchMessages(ctx, in)

	if err != nil {
		if errors.Is(err, &cerror.ValidationError{}) {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, messages)
}

func (mh *MessageHandler) PostMessage(c echo.Context) error {
	var message entities.Message
	if err := c.Bind(&message); err != nil {
		return err
	}
	if ok, err := utility.IsRequestValid(&message); !ok {
		return c.JSON(http.StatusBadRequest, err)
	}

	in := &validation.PostMessageInput{
		Message: &message,
	}

	ctx := c.Request().Context()
	if err := mh.MessageInputPort.PostMessage(ctx, in); err != nil {
		if errors.Is(err, &cerror.ValidationError{}) {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, message)
}

func (mh *MessageHandler) DeleteMessage(c echo.Context) error {
	sMessageID := c.QueryParam("message_id")
	messageID, err := utility.StringToUint(sMessageID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	sUserID := c.QueryParam("user_id")
	userID, err := utility.StringToUint(sUserID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	in := &validation.DeleteMessageInput{
		MessageID: messageID,
		UserID:    userID,
	}

	ctx := c.Request().Context()
	if err := mh.MessageInputPort.DeleteMessage(ctx, in); err != nil {
		if errors.Is(err, &cerror.ValidationError{}) {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func (mh *MessageHandler) UpdateMessage(c echo.Context) error {
	var message entities.Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	in := &validation.UpdateMessageInput{
		Message: &message,
	}

	ctx := c.Request().Context()
	if err := mh.MessageInputPort.UpdateMessage(ctx, in); err != nil {
		if errors.Is(err, &cerror.ValidationError{}) {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}
