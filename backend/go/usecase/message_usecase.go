package usecase

import (
	"context"

	"github.com/shima004/slackclone/model"
	"github.com/shima004/slackclone/repository"
)

type MessageUsecase interface {
	FetchMessages(ctx context.Context, auther string) (res []model.Message, err error)
}

type DefaultMessageUsercase struct {
	MessageRepository repository.MessageRepository
}

func (u *DefaultMessageUsercase) FetchMessages(ctx context.Context, auther string) (res []model.Message, err error) {
	fetchedMessage, _ := u.MessageRepository.FetchMessages(ctx, auther)
	var messages []model.Message
	for _, message := range fetchedMessage {
		messages = append(messages, model.Message{
			UserID: message.UserID,
			Text:   message.Text,
		})
	}
	return messages, nil
}
