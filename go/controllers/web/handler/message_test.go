package handler_test

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
	"github.com/shima004/chat-server/controllers/web/handler"
	"github.com/shima004/chat-server/entities"
	mock_inputport "github.com/shima004/chat-server/mock/inputport"
	"github.com/shima004/chat-server/usecases/inputport/validation"
	"github.com/stretchr/testify/assert"
)

func TestFetchMessages(t *testing.T) {
	t.Parallel()
	mockMessages := []*entities.Message{
		{
			UserID:    453671289,
			ChannelID: 1,
			Text:      "test",
		},
	}

	in := &validation.FatchMessagesInput{
		ChannelID: 1,
		Limit:     1,
		Offset:    0,
	}

	t.Run("FetchMessage", func(t *testing.T) {
		t.Parallel()
		mockctrl := gomock.NewController(t)
		defer mockctrl.Finish()

		mockMessageUsecase := mock_inputport.NewMockMessageUsecase(mockctrl)
		mockMessageUsecase.EXPECT().FetchMessages(gomock.Any(), in).Return(mockMessages, nil).Times(1)

		e := echo.New()
		req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, fmt.Sprintf("/api/messages?channel_id=%d&limit=1&offset=0", mockMessages[0].ChannelID), nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		handler := handler.MessageHandler{MessageInputPort: mockMessageUsecase}
		err = handler.FetchMessages(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var messages []*entities.Message
		err = json.Unmarshal(rec.Body.Bytes(), &messages)
		assert.NoError(t, err)
		assert.Equal(t, mockMessages, messages)
	})
}

func TestPostMessage(t *testing.T) {
	t.Parallel()
	mockMessage := entities.Message{
		UserID:    453671289,
		Text:      "test",
		ChannelID: 1,
	}

	in := &validation.PostMessageInput{
		Message: &mockMessage,
	}

	t.Run("PostMessage", func(t *testing.T) {
		t.Parallel()
		mockctrl := gomock.NewController(t)
		defer mockctrl.Finish()

		mockMessageUsecase := mock_inputport.NewMockMessageUsecase(mockctrl)
		mockMessageUsecase.EXPECT().PostMessage(gomock.Any(), in).Return(nil).Times(1)

		JSON, err := json.Marshal(mockMessage)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, "/api/messages", strings.NewReader(string(JSON)))
		assert.NoError(t, err)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler := handler.MessageHandler{MessageInputPort: mockMessageUsecase}
		err = handler.PostMessage(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)

		var message entities.Message
		err = json.Unmarshal(rec.Body.Bytes(), &message)
		assert.NoError(t, err)
		assert.Equal(t, mockMessage, message)
	})

	t.Run("should return error when invalid json", func(t *testing.T) {
		t.Parallel()
		mockctrl := gomock.NewController(t)
		defer mockctrl.Finish()

		mockMessageUsecase := mock_inputport.NewMockMessageUsecase(mockctrl)
		mockMessageUsecase.EXPECT().PostMessage(gomock.Any(), &mockMessage).Return(nil).Times(0)

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
		handler := handler.MessageHandler{MessageInputPort: mockMessageUsecase}
		err = handler.PostMessage(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
