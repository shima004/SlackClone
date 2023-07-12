package interactor

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shima004/chat-server/entities"
	mock_repository "github.com/shima004/chat-server/mock/repository"
	"github.com/shima004/chat-server/usecases/inputport/validation"
	"github.com/stretchr/testify/assert"
)

func TestCreateChannel(t *testing.T) {
	t.Parallel()

	mockChannel := entities.Channel{
		Name: "test",
	}
	mockChannel.ID = 1

	t.Run("should return nil", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockRepository := mock_repository.NewMockChannelRepository(mockCtrl)
		mockRepository.EXPECT().CreateChannel(gomock.Any(), &mockChannel).Return(mockChannel.ID, nil).Times(1)

		mu := &DefaultChannelUsecase{
			ChannelRepository: mockRepository,
			ContextTimeout:    5,
		}

		in := &validation.CreateChannelInput{
			Channel: &mockChannel,
		}

		id, err := mu.CreateChannel(context.TODO(), in)
		assert.NoError(t, err)
		assert.Equal(t, mockChannel.ID, uint(id))
	})
}

func TestDeleteChannel(t *testing.T) {
	t.Parallel()
	mockChannel := entities.Channel{
		Name: "test",
	}
	mockChannel.ID = 1
	t.Run("should return nil", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockRepository := mock_repository.NewMockChannelRepository(mockCtrl)
		mockRepository.EXPECT().DeleteChannel(gomock.Any(), mockChannel.ID).Return(nil).Times(1)

		mu := &DefaultChannelUsecase{
			ChannelRepository: mockRepository,
			ContextTimeout:    5,
		}

		in := &validation.DeleteChannelInput{
			ChannelID: mockChannel.ID,
		}

		err := mu.DeleteChannel(context.TODO(), in)
		assert.NoError(t, err)
	})
}
