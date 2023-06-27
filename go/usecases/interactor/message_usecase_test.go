package interactor

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shima004/slackclone/entities"
	mock_repository "github.com/shima004/slackclone/mock/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetAllMessages(t *testing.T) {
	mockMessage := entities.Message{
		UserID: 453671289,
		Text:   "Hello World",
	}
	t.Run("should return messagess", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockRepository := mock_repository.NewMockMessageRepository(mockCtrl)
		mockRepository.EXPECT().ReadMessages(gomock.Any(), mockMessage.UserID, 1, 0).Return([]*entities.Message{&mockMessage}, nil).Times(1)

		mu := DefaultMessageUsecase{MessageRepository: mockRepository}
		list, err := mu.FetchMessages(context.TODO(), uint(453671289), 1, 0)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(list))
		assert.Equal(t, mockMessage.UserID, list[0].UserID)
		assert.Equal(t, mockMessage.Text, list[0].Text)
	})

	t.Run("should return error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockRepository := mock_repository.NewMockMessageRepository(mockCtrl)
		mockRepository.EXPECT().ReadMessages(gomock.Any(), mockMessage.UserID, 1, 0).Return(nil, errors.New("Unexpected Error")).Times(1)

		mu := DefaultMessageUsecase{MessageRepository: mockRepository}
		list, err := mu.FetchMessages(context.TODO(), uint(453671289), 1, 0)
		assert.Error(t, err)
		assert.Nil(t, list)
	})
}

func TestPostMessage(t *testing.T) {
	mockMessage := entities.Message{
		UserID:    453671289,
		Text:      "Hello World",
		ChannelID: 1,
	}
	t.Run("should return nil", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockMRepository := mock_repository.NewMockMessageRepository(mockCtrl)
		mockMRepository.EXPECT().CreateMessage(gomock.Any(), &mockMessage).Return(uint(1), nil).Times(1)

		mockCRepository := mock_repository.NewMockChannelRepository(mockCtrl)
		mockCRepository.EXPECT().ReadChannel(gomock.Any(), mockMessage.ChannelID).Return(&entities.Channel{}, nil).Times(1)

		mu := DefaultMessageUsecase{MessageRepository: mockMRepository, ChannelRepository: mockCRepository}
		err := mu.PostMessage(context.Background(), &mockMessage)
		assert.NoError(t, err)
	})
}

func TestDeleteMessage(t *testing.T) {
	mockMessage := entities.Message{
		UserID: 453671289,
		Text:   "Hello World",
	}
	t.Run("should return nil", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockRepository := mock_repository.NewMockMessageRepository(mockCtrl)
		mockRepository.EXPECT().DeleteMessage(gomock.Any(), mockMessage.ID).Return(nil).Times(1)

		mu := DefaultMessageUsecase{MessageRepository: mockRepository}
		err := mu.DeleteMessage(context.Background(), mockMessage.ID)
		assert.NoError(t, err)
	})
}

func TestUpdateMessage(t *testing.T) {
	mockMessage := entities.Message{
		UserID: 453671289,
		Text:   "Hello World",
	}
	t.Run("should return nil", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockRepository := mock_repository.NewMockMessageRepository(mockCtrl)
		mockRepository.EXPECT().UpdateMessage(gomock.Any(), &mockMessage).Return(nil).Times(1)

		mu := DefaultMessageUsecase{MessageRepository: mockRepository}
		err := mu.UpdateMessage(context.Background(), &mockMessage)
		assert.NoError(t, err)
	})
}
