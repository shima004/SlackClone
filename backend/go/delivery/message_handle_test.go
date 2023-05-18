package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/shima004/slackclone/model"
	"github.com/shima004/slackclone/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchMessages(t *testing.T) {
	t.Run("FetchMessage", func(t *testing.T) {
		mockMessageUsecase := new(usecase.MockMessageUsercase)
		mockMessages := []model.Message{
			{
				UserID: 453671289,
				Text:   "test",
			},
		}
		mockMessageUsecase.On("FetchMessages", mock.Anything, mock.AnythingOfType("uint")).Return(mockMessages, nil)
		e := echo.New()
		req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, fmt.Sprintf("/api/messages?user_id=%d", mockMessages[0].UserID), nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		handler := MessageHandler{
			MessageUseCase: mockMessageUsecase,
		}
		err = handler.FetchMessages(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var messages []model.Message
		err = json.Unmarshal(rec.Body.Bytes(), &messages)
		assert.NoError(t, err)
		assert.Equal(t, mockMessages, messages)

		mockMessageUsecase.AssertExpectations(t)
	})
}

func TestPostMessage(t *testing.T) {
	t.Run("PostMessage", func(t *testing.T) {
		mockMessageUsecase := new(usecase.MockMessageUsercase)
		mockMessage := model.Message{
			UserID:    453671289,
			Text:      "test",
			ChannelID: 1,
		}
		JSON, err := json.Marshal(mockMessage)
		assert.NoError(t, err)

		mockMessageUsecase.On("PostMessage", mock.Anything, mock.AnythingOfType("model.Message")).Return(nil).Once()
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
		assert.Equal(t, http.StatusOK, rec.Code)

		var message model.Message
		err = json.Unmarshal(rec.Body.Bytes(), &message)
		assert.NoError(t, err)
		assert.Equal(t, mockMessage, message)

		mockMessageUsecase.AssertExpectations(t)
	})

	t.Run("should return error when invalid json", func(t *testing.T) {
		mockMessageUsecase := new(usecase.MockMessageUsercase)
		mockMessage := model.Message{
			UserID:    453671289,
			ChannelID: 1,
		}
		JSON, err := json.Marshal(mockMessage)
		assert.NoError(t, err)

		mockMessageUsecase.On("PostMessage", mock.Anything, mock.AnythingOfType("model.Message")).Return(nil)
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
