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
	mockRepository := repository.MockRepository{}
	mockMessage := model.Message{
		Auther: "pacapaca",
		Text:   "Hello World",
	}
	t.Run("should return messagess", func(t *testing.T) {
		mockRepository.On("FetchMessages", mock.Anything, mock.AnythingOfType("string")).Return([]model.Message{
			mockMessage,
		}, nil).Once()

		mu := MessageUsercase{&mockRepository}
		list, err := mu.FetchMessages(context.TODO(), "pacapaca")
		assert.NoError(t, err)
		assert.Equal(t, 1, len(list))
		assert.Equal(t, mockMessage.Auther, list[0].Auther)
		assert.Equal(t, mockMessage.Text, list[0].Text)
	})
}
