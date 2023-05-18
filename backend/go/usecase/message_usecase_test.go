package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/shima004/slackclone/model"
	"github.com/shima004/slackclone/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllMessages(t *testing.T) {
	mockRepository := repository.MockMessageRepository{}
	mockMessage := model.Message{
		UserID: 453671289,
		Text:   "Hello World",
	}
	t.Run("should return messagess", func(t *testing.T) {
		mockRepository.On("FetchMessages", mock.Anything, mock.AnythingOfType("uint")).Return([]model.Message{
			mockMessage,
		}, nil).Once()

		mu := DefaultMessageUsecase{MessageRepository: &mockRepository}
		list, err := mu.FetchMessages(context.TODO(), uint(453671289))
		assert.NoError(t, err)
		assert.Equal(t, 1, len(list))
		assert.Equal(t, mockMessage.UserID, list[0].UserID)
		assert.Equal(t, mockMessage.Text, list[0].Text)
	})

	t.Run("should return error", func(t *testing.T) {
		mockRepository.On("FetchMessages", mock.Anything, mock.AnythingOfType("uint")).Return(nil, errors.New("Unexpected Error")).Once()

		mu := DefaultMessageUsecase{MessageRepository: &mockRepository}
		list, err := mu.FetchMessages(context.TODO(), uint(453671289))
		assert.Error(t, err)
		assert.Nil(t, list)
	})
}
