package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/shima004/slackclone/model"
	"github.com/stretchr/testify/assert"
)

func TestGetAllMessages(t *testing.T) {
	t.Run("should return 200 status ok", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/messages", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		controller := Controller{}
		controller.GetAllMessages(c)
		var messages []model.Message
		json.Unmarshal(rec.Body.Bytes(), &messages)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, 1, len(messages))
		assert.Equal(t, "PacaPaca", messages[0].Auther)
		assert.Equal(t, "Hello World", messages[0].Text)
	})
}
