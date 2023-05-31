//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

package usecase

import (
	"context"
	"time"

	"github.com/shima004/slackclone/model"
	"github.com/shima004/slackclone/repository"
)

type MessageUsecase interface {
	FetchMessages(ctx context.Context, channelID uint, limit int, offset int) (res []model.Message, err error)
	PostMessage(ctx context.Context, message model.Message) (err error)
	DeleteMessage(ctx context.Context, messageID uint) (err error)
	UpdateMessage(ctx context.Context, message model.Message) (err error)
}

type DefaultMessageUsecase struct {
	MessageRepository repository.MessageRepository
	ChannelRepository repository.ChannelRepository
	ContextTimeout    time.Duration
}

func (u *DefaultMessageUsecase) FetchMessages(ctx context.Context, channelID uint, limit int, offset int) (res []model.Message, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	fetchedMessage, err := u.MessageRepository.FetchMessages(ctx, channelID, limit, offset)

	if err != nil {
		return nil, err
	}

	return fetchedMessage, nil
}

func (u *DefaultMessageUsecase) PostMessage(ctx context.Context, message model.Message) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	// Check if channel exists
	_, err = u.ChannelRepository.FetchChannel(ctx, message.ChannelID)
	if err != nil {
		return err
	}

	err = u.MessageRepository.PostMessage(ctx, message)
	return err
}

func (u *DefaultMessageUsecase) DeleteMessage(ctx context.Context, messageID uint) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	err = u.MessageRepository.DeleteMessage(ctx, messageID)
	return err
}

func (u *DefaultMessageUsecase) UpdateMessage(ctx context.Context, message model.Message) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	err = u.MessageRepository.UpdateMessage(ctx, message)
	return err
}
