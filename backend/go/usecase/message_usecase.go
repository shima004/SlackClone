package usecase

import (
	"context"

	"github.com/shima004/slackclone/model"
	"github.com/shima004/slackclone/repository"
)

type MessageUsercase struct {
	repository.Repository
}

func (u *MessageUsercase) FetchMessages(ctx context.Context, auther string) (res []model.Message, err error) {
	fetchedMessage, _ := u.Repository.FetchMessages(ctx, auther)
	var messages []model.Message
	for _, message := range fetchedMessage {
		messages = append(messages, model.Message{
			Auther: message.Auther,
			Text:   message.Text,
		})
	}
	return messages, nil
}
