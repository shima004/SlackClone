//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

package usecase

import (
	"context"
	"time"

	"github.com/shima004/slackclone/model"
	"github.com/shima004/slackclone/repository"
)

type ChannelUsecase interface {
	CreateChannel(ctx context.Context, channel *model.Channel) (uint, error)
	DeleteChannel(ctx context.Context, channelID uint) error
}

type DefaultChannelUsecase struct {
	ChannelRepository repository.ChannelRepository
	contextTimeout    time.Duration
}

func (u *DefaultChannelUsecase) CreateChannel(ctx context.Context, channel *model.Channel) (uint, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	channelID, err := u.ChannelRepository.CreateChannel(ctx, channel)
	return channelID, err
}

func (u *DefaultChannelUsecase) DeleteChannel(ctx context.Context, channelID uint) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	err := u.ChannelRepository.DeleteChannel(ctx, channelID)
	return err
}
