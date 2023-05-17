package delivery

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
		mockMessageUsecase.On("FetchMessages", mock.Anything, mock.AnythingOfType("string")).Return(mockMessages, nil)
		e := echo.New()
		req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, "/api/messages", nil)
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
