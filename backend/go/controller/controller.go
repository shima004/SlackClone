package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/shima004/slackclone/model"
	"github.com/shima004/slackclone/repository"
)

type Controller struct {
	repository.Repository
}

func (m *Controller) GetAllMessages(c echo.Context) error {
	messages := []model.Message{
		{ID: 1, Auther: "PacaPaca", Text: "Hello World"},
	}
	return c.JSON(200, messages)
}
