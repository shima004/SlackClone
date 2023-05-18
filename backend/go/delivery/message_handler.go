package delivery

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
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
}

func (mh *MessageHandler) FetchMessages(c echo.Context) error {
	sUserID := c.QueryParam("user_id")
	userID, err := StringToUint(sUserID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	messages, err := mh.MessageUseCase.FetchMessages(ctx, userID)
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

func isRequestValid(m *model.Message) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func StringToUint(s string) (uint, error) {
	u, err := strconv.ParseUint(s, 10, 32)
	return uint(u), err
}
