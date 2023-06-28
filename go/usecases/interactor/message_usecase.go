//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

package interactor

import (
	"context"
	"time"

	"github.com/shima004/chat-server/entities"
	"github.com/shima004/chat-server/usecases/repository"
)

type DefaultMessageUsecase struct {
	MessageRepository repository.MessageRepository
	ChannelRepository repository.ChannelRepository
	ContextTimeout    time.Duration
}

func (u *DefaultMessageUsecase) FetchMessages(ctx context.Context, channelID uint, limit int, offset int) (res []*entities.Message, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	fetchedMessage, err := u.MessageRepository.ReadMessages(ctx, channelID, limit, offset)

	if err != nil {
		return nil, err
	}

	return fetchedMessage, nil
}

func (u *DefaultMessageUsecase) PostMessage(ctx context.Context, message *entities.Message) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	// Check if channel exists
	_, err = u.ChannelRepository.ReadChannel(ctx, message.ChannelID)
	if err != nil {
		return err
	}

	_, err = u.MessageRepository.CreateMessage(ctx, message)
	return err
}

func (u *DefaultMessageUsecase) DeleteMessage(ctx context.Context, messageID uint, userID uint) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	fetchedMessage, err := u.MessageRepository.ReadMessage(ctx, messageID)
	if err != nil {
		return err
	}
	if fetchedMessage.UserID != userID {
		return entities.ErrUnauthorized
	}

	err = u.MessageRepository.DeleteMessage(ctx, messageID)
	return err
}

func (u *DefaultMessageUsecase) UpdateMessage(ctx context.Context, message *entities.Message) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	err = u.MessageRepository.UpdateMessage(ctx, message)
	return err
}
