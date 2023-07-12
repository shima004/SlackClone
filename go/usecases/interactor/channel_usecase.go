package interactor

import (
	"context"
	"time"

	"github.com/shima004/chat-server/entities"
	"github.com/shima004/chat-server/usecases/inputport/validation"
	"github.com/shima004/chat-server/usecases/repository"
)

type DefaultChannelUsecase struct {
	ChannelRepository repository.ChannelRepository
	ContextTimeout    time.Duration
}

func (u *DefaultChannelUsecase) CreateChannel(ctx context.Context, in *validation.CreateChannelInput) (uint, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	err := in.Validate()
	if err != nil {
		return 0, err
	}

	channelID, err := u.ChannelRepository.CreateChannel(ctx, in.Channel)
	return channelID, err
}

func (u *DefaultChannelUsecase) DeleteChannel(ctx context.Context, in *validation.DeleteChannelInput) error {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	err := in.Validate()
	if err != nil {
		return err
	}

	err = u.ChannelRepository.DeleteChannel(ctx, in.ChannelID)
	return err
}

func (u *DefaultChannelUsecase) FetchChannel(ctx context.Context, in *validation.FetchChannelInput) (*entities.Channel, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	err := in.Validate()
	if err != nil {
		return nil, err
	}

	channel, err := u.ChannelRepository.ReadChannel(ctx, in.ChannelID)
	return channel, err
}
