package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shima004/chat-server/entities"
	"github.com/shima004/chat-server/usecases/inputport"
	"github.com/shima004/chat-server/usecases/inputport/validation"
	"github.com/shima004/chat-server/utility"
)

type ChannelHandler struct {
	ChannelInputPort inputport.ChannelUsecase
}

func NewChannelHandler(cip inputport.ChannelUsecase) *ChannelHandler {
	return &ChannelHandler{cip}
}

func (ch *ChannelHandler) PostChannel(c echo.Context) error {
	var channel *entities.Channel
	if err := c.Bind(&channel); err != nil {
		return err
	}

	ctx := c.Request().Context()
	in := &validation.CreateChannelInput{Channel: channel}

	channelID, err := ch.ChannelInputPort.CreateChannel(ctx, in)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, channelID)
}

func (ch *ChannelHandler) DeleteChannel(c echo.Context) error {
	sChannelID := c.Param("channelID")
	channelID, err := utility.StringToUint(sChannelID)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	in := &validation.DeleteChannelInput{ChannelID: channelID}

	if err := ch.ChannelInputPort.DeleteChannel(ctx, in); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (ch *ChannelHandler) FetchChannel(c echo.Context) error {
	sChannelID := c.Param("channelID")
	channelID, err := utility.StringToUint(sChannelID)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	in := &validation.FetchChannelInput{ChannelID: channelID}

	channel, err := ch.ChannelInputPort.FetchChannel(ctx, in)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, channel)
}
