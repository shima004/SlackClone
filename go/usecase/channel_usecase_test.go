package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mock_repository "github.com/shima004/slackclone/mock/repository"
	"github.com/shima004/slackclone/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateChannel(t *testing.T) {
	mockChannel := model.Channel{
		Name: "test",
	}
	mockChannel.ID = 1
	t.Run("should return nil", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockRepository := mock_repository.NewMockChannelRepository(mockCtrl)
		mockRepository.EXPECT().CreateChannel(gomock.Any(), &mockChannel).Return(mockChannel.ID, nil).Times(1)

		mu := &DefaultChannelUsecase{
			ChannelRepository: mockRepository,
			ContextTimeout:    5,
		}

		id, err := mu.CreateChannel(context.TODO(), &mockChannel)
		assert.NoError(t, err)
		assert.Equal(t, mockChannel.ID, uint(id))
	})
}

func TestDeleteChannel(t *testing.T) {
	mockChannel := model.Channel{
		Name: "test",
	}
	mockChannel.ID = 1
	t.Run("should return nil", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockRepository := mock_repository.NewMockChannelRepository(mockCtrl)
		mockRepository.EXPECT().DeleteChannel(gomock.Any(), mockChannel.ID).Return(nil).Times(1)

		mu := &DefaultChannelUsecase{
			ChannelRepository: mockRepository,
			ContextTimeout:    5,
		}

		err := mu.DeleteChannel(context.TODO(), mockChannel.ID)
		assert.NoError(t, err)
	})
}
