package usecase

import (
	"context"
	"time"

	"github.com/shima004/slackclone/model"
	"github.com/shima004/slackclone/repository"
)

type MessageUsecase interface {
	FetchMessages(ctx context.Context, userID uint) (res []model.Message, err error)
	PostMessage(ctx context.Context, message model.Message) (err error)
}

type DefaultMessageUsecase struct {
	MessageRepository repository.MessageRepository
	contextTimeout    time.Duration
}

func (u *DefaultMessageUsecase) FetchMessages(ctx context.Context, userID uint) (res []model.Message, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	fetchedMessage, err := u.MessageRepository.FetchMessages(ctx, userID)

	if err != nil {
		return nil, err
	}

	var messages []model.Message
	for _, message := range fetchedMessage {
		messages = append(messages, model.Message{
			UserID: message.UserID,
			Text:   message.Text,
		})
	}
	return messages, nil
}

func (u *DefaultMessageUsecase) PostMessage(ctx context.Context, message model.Message) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	// message.CreatedAt = time.Now()
	// message.UpdatedAt = time.Now()
	
	err = u.MessageRepository.PostMessage(ctx, message)
	return err
}
