package usecase

import (
	"context"
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
		mockRepository.On("FetchMessages", mock.Anything, mock.AnythingOfType("string")).Return([]model.Message{
			mockMessage,
		}, nil).Once()

		mu := DefaultMessageUsercase{MessageRepository: &mockRepository}
		list, err := mu.FetchMessages(context.TODO(), "pacapaca")
		assert.NoError(t, err)
		assert.Equal(t, 1, len(list))
		assert.Equal(t, mockMessage.UserID, list[0].UserID)
		assert.Equal(t, mockMessage.Text, list[0].Text)
	})
}
