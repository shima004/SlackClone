package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shima004/slackclone/entities"
	"github.com/shima004/slackclone/usecase"
)

type ChannelHandler struct {
	ChannelUseCase usecase.ChannelUsecase
}

func NewChannelHandler(g *echo.Group, cu usecase.ChannelUsecase) {
	handler := &ChannelHandler{cu}
	g.POST("/channels", handler.CreateChannel)
	g.DELETE("/channels/:channelID", handler.DeleteChannel)
}

func (ch *ChannelHandler) CreateChannel(c echo.Context) error {
	var channel entities.Channel
	if err := c.Bind(&channel); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if ok, err := isRequestValid(channel); !ok {
		return c.JSON(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	channelID, err := ch.ChannelUseCase.CreateChannel(ctx, &channel)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusCreated, channelID)
}

func (ch *ChannelHandler) DeleteChannel(c echo.Context) error {
	sChannelID := c.Param("channelID")
	channelID, err := StringToUint(sChannelID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	err = ch.ChannelUseCase.DeleteChannel(ctx, channelID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, nil)
}
