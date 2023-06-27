//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

package interactor

import (
	"context"
	"time"

	"github.com/shima004/chat-server/entities"
	"github.com/shima004/chat-server/usecases/repository"
)

type DefaultChannelUsecase struct {
	ChannelRepository repository.ChannelRepository
	ContextTimeout    time.Duration
}

func (u *DefaultChannelUsecase) CreateChannel(ctx context.Context, channel *entities.Channel) (uint, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	channelID, err := u.ChannelRepository.CreateChannel(ctx, channel)
	return channelID, err
}

func (u *DefaultChannelUsecase) DeleteChannel(ctx context.Context, channelID uint) error {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	err := u.ChannelRepository.DeleteChannel(ctx, channelID)
	return err
}

func (u *DefaultChannelUsecase) FetchChannel(ctx context.Context, channelID uint) (*entities.Channel, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	channel, err := u.ChannelRepository.ReadChannel(ctx, channelID)
	return channel, err
}
