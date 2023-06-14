package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/shima004/slackclone/entities"
	mock_usecase "github.com/shima004/slackclone/mock/usecase"
	"github.com/stretchr/testify/assert"
)

func TestFetchMessages(t *testing.T) {
	mockMessages := []entities.Message{
		{
			UserID:    453671289,
			ChannelID: 1,
			Text:      "test",
		},
	}
	t.Run("FetchMessage", func(t *testing.T) {
		mockctrl := gomock.NewController(t)
		defer mockctrl.Finish()

		mockMessageUsecase := mock_usecase.NewMockMessageUsecase(mockctrl)
		mockMessageUsecase.EXPECT().FetchMessages(gomock.Any(), mockMessages[0].ChannelID, 1, 0).Return(mockMessages, nil).Times(1)

		e := echo.New()
		req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, fmt.Sprintf("/api/messages?channel_id=%d&limit=1&offset=0", mockMessages[0].ChannelID), nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		handler := MessageHandler{
			MessageUseCase: mockMessageUsecase,
		}
		err = handler.FetchMessages(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var messages []entities.Message
		err = json.Unmarshal(rec.Body.Bytes(), &messages)
		assert.NoError(t, err)
		assert.Equal(t, mockMessages, messages)
	})
}

func TestPostMessage(t *testing.T) {
	mockMessage := entities.Message{
		UserID:    453671289,
		Text:      "test",
		ChannelID: 1,
	}
	t.Run("PostMessage", func(t *testing.T) {
		mockctrl := gomock.NewController(t)
		defer mockctrl.Finish()

		mockMessageUsecase := mock_usecase.NewMockMessageUsecase(mockctrl)
		mockMessageUsecase.EXPECT().PostMessage(gomock.Any(), mockMessage).Return(nil).Times(1)

		JSON, err := json.Marshal(mockMessage)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, "/api/messages", strings.NewReader(string(JSON)))
		assert.NoError(t, err)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler := MessageHandler{
			MessageUseCase: mockMessageUsecase,
		}
		err = handler.PostMessage(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)

		var message entities.Message
		err = json.Unmarshal(rec.Body.Bytes(), &message)
		assert.NoError(t, err)
		assert.Equal(t, mockMessage, message)
	})

	t.Run("should return error when invalid json", func(t *testing.T) {
		mockctrl := gomock.NewController(t)
		defer mockctrl.Finish()

		mockMessageUsecase := mock_usecase.NewMockMessageUsecase(mockctrl)
		mockMessageUsecase.EXPECT().PostMessage(gomock.Any(), mockMessage).Return(nil).Times(0)

		mockMessage := entities.Message{
			UserID:    453671289,
			ChannelID: 1,
		}
		JSON, err := json.Marshal(mockMessage)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, "/api/messages", strings.NewReader(string(JSON)))
		assert.NoError(t, err)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler := MessageHandler{
			MessageUseCase: mockMessageUsecase,
		}

		err = handler.PostMessage(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
