package repository

import (
	"context"

	"github.com/shima004/slackclone/model"
	"gorm.io/gorm"
)

type MessageRepository interface {
	FetchMessages(context.Context, string) ([]model.Message, error)
	PostMessage(context.Context, model.Message) error
}

type DefaultMessageRepository struct {
}

func (r *DefaultMessageRepository) FetchMessages(ctx context.Context, auther string) ([]model.Message, error) {
	return []model.Message{
		{UserID: 453671289, Text: "Hello World", Model: gorm.Model{ID: 1}},
	}, nil
}

func (r *DefaultMessageRepository) PostMessage(ctx context.Context, message model.Message) error {
	return nil
}
