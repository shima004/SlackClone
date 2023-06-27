package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shima004/slackclone/entities"
	"github.com/shima004/slackclone/usecases/inputport"
	"github.com/shima004/slackclone/utility"
)

type ChannelHandler struct {
	ChannelInputPort inputport.ChannelUsecase
}

func NewChannelHandler(cip inputport.ChannelUsecase) *ChannelHandler {
	return &ChannelHandler{cip}
}

func (ch *ChannelHandler) PostChannel(c echo.Context) error {
	var channel entities.Channel
	if err := c.Bind(&channel); err != nil {
		return err
	}
	ctx := c.Request().Context()
	channelID, err := ch.ChannelInputPort.CreateChannel(ctx, &channel)
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
	if err := ch.ChannelInputPort.DeleteChannel(ctx, channelID); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
