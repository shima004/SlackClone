package interactor

import (
	"context"
	"time"

	"github.com/shima004/chat-server/entities"
	"github.com/shima004/chat-server/usecases/inputport/validation"
	"github.com/shima004/chat-server/usecases/repository"
)

type DefaultMessageUsecase struct {
	MessageRepository repository.MessageRepository
	ChannelRepository repository.ChannelRepository
	ContextTimeout    time.Duration
}

func (u *DefaultMessageUsecase) FetchMessages(ctx context.Context, in *validation.FatchMessagesInput) (res []*entities.Message, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	err = in.Validate()
	if err != nil {
		return nil, err
	}

	channelID := in.ChannelID
	limit := in.Limit
	offset := in.Offset
	fetchedMessage, err := u.MessageRepository.ReadMessages(ctx, channelID, limit, offset)
	if err != nil {
		return nil, err
	}

	return fetchedMessage, nil
}

func (u *DefaultMessageUsecase) PostMessage(ctx context.Context, in *validation.PostMessageInput) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	err = in.Validate()
	if err != nil {
		return err
	}

	message := in.Message
	_, err = u.ChannelRepository.ReadChannel(ctx, message.ChannelID)
	if err != nil {
		return err
	}

	_, err = u.MessageRepository.CreateMessage(ctx, message)
	return err
}

func (u *DefaultMessageUsecase) DeleteMessage(ctx context.Context, in *validation.DeleteMessageInput) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	err = in.Validate()
	if err != nil {
		return err
	}

	messageID := in.MessageID
	userID := in.UserID

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

func (u *DefaultMessageUsecase) UpdateMessage(ctx context.Context, in *validation.UpdateMessageInput) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	err = in.Validate()
	if err != nil {
		return err
	}

	message := in.Message

	err = u.MessageRepository.UpdateMessage(ctx, message)
	return err
}
