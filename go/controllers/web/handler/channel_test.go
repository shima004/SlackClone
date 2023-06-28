package handler_test

// import (
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/labstack/echo/v4"
// 	"github.com/shima004/chat-server/controllers/web/handler"
// 	"github.com/shima004/chat-server/entities"
// 	mock_inputport "github.com/shima004/chat-server/mock/inputport"
// 	"github.com/stretchr/testify/assert"
// )

// func TestPostChannel(t *testing.T) {
// 	t.Parallel()
// 	mockChannel := entities.Channel{
// 		Name: "test",
// 	}

// 	t.Run("should return nil", func(t *testing.T) {
// 		t.Parallel()
// 		mockCtrl := gomock.NewController(t)
// 		defer mockCtrl.Finish()

// 		mockChannelUsecase := mock_inputport.NewMockChannelUsecase(mockCtrl)
// 		mockChannelUsecase.EXPECT().CreateChannel(gomock.Any(), &mockChannel).Return(uint(1), nil).Times(1)
// 		JSON, err := json.Marshal(mockChannel)
// 		assert.NoError(t, err)

// 		e := echo.New()
// 		req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, "/api/channels", strings.NewReader(string(JSON)))
// 		assert.NoError(t, err)
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		rec := httptest.NewRecorder()
// 		c := e.NewContext(req, rec)
// 		handler := handler.ChannelHandler{
// 			ChannelInputPort: mockChannelUsecase,
// 		}
// 		err = handler.PostChannel(c)
// 		assert.NoError(t, err)
// 		assert.Equal(t, http.StatusCreated, rec.Code)

// 		var channelID uint
// 		err = json.Unmarshal(rec.Body.Bytes(), &channelID)
// 		assert.NoError(t, err)
// 		assert.Equal(t, uint(1), channelID)
// 	})
// 	t.Run("should return StatusUnprocessableEntity", func(t *testing.T) {
// 		t.Parallel()
// 		mockCtrl := gomock.NewController(t)
// 		defer mockCtrl.Finish()

// 		mockChannelUsecase := mock_inputport.NewMockChannelUsecase(mockCtrl)
// 		mockChannelUsecase.EXPECT().CreateChannel(gomock.Any(), &mockChannel).Return(uint(1), nil).Times(0)

// 		e := echo.New()
// 		req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, "/api/channels", strings.NewReader("invalid json"))
// 		assert.NoError(t, err)
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		rec := httptest.NewRecorder()
// 		c := e.NewContext(req, rec)
// 		handler := handler.ChannelHandler{
// 			ChannelInputPort: mockChannelUsecase,
// 		}
// 		err = handler.PostChannel(c)
// 		assert.NoError(t, err)
// 		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
// 	})
// 	t.Run("should return StatusBadRequest", func(t *testing.T) {
// 		t.Parallel()
// 		mockCtrl := gomock.NewController(t)
// 		defer mockCtrl.Finish()

// 		mockChannelUsecase := mock_inputport.NewMockChannelUsecase(mockCtrl)
// 		mockChannelUsecase.EXPECT().CreateChannel(gomock.Any(), &mockChannel).Return(uint(1), nil).Times(0)
// 		JSON, err := json.Marshal(entities.Channel{})
// 		assert.NoError(t, err)

// 		e := echo.New()
// 		req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, "/api/channels", strings.NewReader(string(JSON)))
// 		assert.NoError(t, err)
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		rec := httptest.NewRecorder()
// 		c := e.NewContext(req, rec)
// 		handler := handler.ChannelHandler{
// 			ChannelInputPort: mockChannelUsecase,
// 		}
// 		err = handler.PostChannel(c)
// 		assert.NoError(t, err)
// 		assert.Equal(t, http.StatusBadRequest, rec.Code)
// 	})
// 	t.Run("should return StatusNotFound", func(t *testing.T) {
// 		t.Parallel()
// 		mockCtrl := gomock.NewController(t)
// 		defer mockCtrl.Finish()

// 		mockChannelUsecase := mock_inputport.NewMockChannelUsecase(mockCtrl)
// 		mockChannelUsecase.EXPECT().CreateChannel(gomock.Any(), &mockChannel).Return(uint(0), errors.New("not found")).Times(1)
// 		JSON, err := json.Marshal(mockChannel)
// 		assert.NoError(t, err)

// 		e := echo.New()
// 		req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, "/api/channels", strings.NewReader(string(JSON)))
// 		assert.NoError(t, err)
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		rec := httptest.NewRecorder()
// 		c := e.NewContext(req, rec)
// 		handler := handler.ChannelHandler{
// 			ChannelInputPort: mockChannelUsecase,
// 		}
// 		err = handler.PostChannel(c)
// 		assert.NoError(t, err)
// 		assert.Equal(t, http.StatusNotFound, rec.Code)
// 	})
// }

// func TestDeleteChannel(t *testing.T) {
// 	t.Parallel()
// 	t.Run("should return nil", func(t *testing.T) {
// 		t.Parallel()
// 		mockCtrl := gomock.NewController(t)
// 		defer mockCtrl.Finish()

// 		mockChannelUsecase := mock_inputport.NewMockChannelUsecase(mockCtrl)
// 		mockChannelUsecase.EXPECT().DeleteChannel(gomock.Any(), uint(1)).Return(nil).Times(1)

// 		e := echo.New()
// 		req, err := http.NewRequestWithContext(context.TODO(), http.MethodDelete, "/api/channels/1", nil)
// 		assert.NoError(t, err)

// 		rec := httptest.NewRecorder()
// 		c := e.NewContext(req, rec)
// 		c.SetParamNames("channelID")
// 		c.SetParamValues("1")
// 		handler := handler.ChannelHandler{
// 			ChannelInputPort: mockChannelUsecase,
// 		}
// 		err = handler.DeleteChannel(c)
// 		assert.NoError(t, err)
// 		assert.Equal(t, http.StatusNoContent, rec.Code)
// 	})

// 	t.Run("should return StatusBadRequest", func(t *testing.T) {
// 		t.Parallel()
// 		mockCtrl := gomock.NewController(t)
// 		defer mockCtrl.Finish()

// 		mockChannelUsecase := mock_inputport.NewMockChannelUsecase(mockCtrl)
// 		mockChannelUsecase.EXPECT().DeleteChannel(gomock.Any(), uint(1)).Return(nil).Times(0)

// 		e := echo.New()
// 		req, err := http.NewRequestWithContext(context.TODO(), http.MethodDelete, "/api/channels/invalid", nil)
// 		assert.NoError(t, err)

// 		rec := httptest.NewRecorder()
// 		c := e.NewContext(req, rec)
// 		c.SetParamNames("channelID")
// 		c.SetParamValues("invalid")
// 		handler := handler.ChannelHandler{
// 			ChannelInputPort: mockChannelUsecase,
// 		}
// 		err = handler.DeleteChannel(c)
// 		assert.NoError(t, err)
// 		assert.Equal(t, http.StatusBadRequest, rec.Code)
// 	})

// 	t.Run("should return StatusNotFound", func(t *testing.T) {
// 		t.Parallel()
// 		mockCtrl := gomock.NewController(t)
// 		defer mockCtrl.Finish()

// 		mockChannelUsecase := mock_inputport.NewMockChannelUsecase(mockCtrl)
// 		mockChannelUsecase.EXPECT().DeleteChannel(gomock.Any(), uint(1)).Return(errors.New("not found")).Times(1)

// 		e := echo.New()
// 		req, err := http.NewRequestWithContext(context.TODO(), http.MethodDelete, "/api/channels/1", nil)
// 		assert.NoError(t, err)

// 		rec := httptest.NewRecorder()
// 		c := e.NewContext(req, rec)
// 		c.SetParamNames("channelID")
// 		c.SetParamValues("1")
// 		handler := handler.ChannelHandler{
// 			ChannelInputPort: mockChannelUsecase,
// 		}
// 		err = handler.DeleteChannel(c)
// 		assert.NoError(t, err)
// 		assert.Equal(t, http.StatusNotFound, rec.Code)
// 	})
// }
