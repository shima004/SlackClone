package repository

import (
	"context"

	"github.com/shima004/slackclone/model"
	"gorm.io/gorm"
)

type Repository interface {
	FetchMessages(context.Context, string) ([]model.Message, error)
}

type DefaultRepository struct {
}

func (r *DefaultRepository) FetchMessages(ctx context.Context, auther string) ([]model.Message, error) {
	return []model.Message{
		{Auther: "pacapaca", Text: "Hello World", Model: gorm.Model{ID: 1}},
	}, nil
}
